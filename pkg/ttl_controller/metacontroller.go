package ttl_controller

import (
	"context"
	"fmt"
	"github.com/fpetkovski/k8s-ttl-controller/pkg/apis/fpetkovski_io/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizer = "fpetkovski.io/ttl-controller"
)

// metacontroller reconciles TTL Controllers
type metacontroller struct {
	client            client.Client
	manager           controllerruntime.Manager
	parentControllers map[string]*ttlController
	logger            logr.Logger
}

func NewMetacontroller(mgr controllerruntime.Manager, logger logr.Logger) *metacontroller {
	return &metacontroller{
		client:            mgr.GetClient(),
		manager:           mgr,
		parentControllers: make(map[string]*ttlController),
		logger:            logger.WithName("metacontroller"),
	}
}

// Implement reconcile.Reconciler so the controller can reconcile objects
var _ reconcile.Reconciler = &metacontroller{}

func (mc *metacontroller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	cc := &v1alpha1.TTLPolicy{}
	if err := mc.client.Get(context.TODO(), request.NamespacedName, cc); errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not get TTLPolicy: %+v", err)
	}

	parentCtrlName := request.NamespacedName.String()
	if cc.GetDeletionTimestamp() != nil {
		mc.logger.Info("TTLPolicy deleted, removing controller", "Name", request.NamespacedName)
		return mc.deleteController(parentCtrlName, cc)
	}

	if !controllerutil.ContainsFinalizer(cc, finalizer) {
		mc.logger.Info("Setting finalizer on TTLPolicy", "Name", request.Name)
		controllerutil.AddFinalizer(cc, finalizer)
		if err := mc.client.Update(context.TODO(), cc); err != nil {
			return reconcile.Result{}, fmt.Errorf("could not set finalizer to TTLPolicy: %+v", err)
		}
		return reconcile.Result{}, nil
	}

	return mc.syncController(parentCtrlName, cc)
}

func (mc *metacontroller) syncController(ctrlName string, cc *v1alpha1.TTLPolicy) (reconcile.Result, error) {
	if ctrl, found := mc.parentControllers[ctrlName]; found {
		mc.logger.Info("TTL controller already exists, stopping", "Name", ctrlName)
		ctrl.Stop()
		delete(mc.parentControllers, ctrlName)
	}

	resourceSpec := cc.Spec.Resource
	groupVersion, err := schema.ParseGroupVersion(resourceSpec.APIVersion)
	if err != nil {
		mc.logger.Error(err, "Could not create new ttl controller")
		return reconcile.Result{}, nil
	}

	gvk := groupVersion.WithKind(resourceSpec.Kind)
	ctrl, err := newTTLController(ctrlName, mc.manager, gvk, cc.Spec.TTLFrom, cc.Spec.ExpirationFrom)
	if err != nil {
		mc.logger.Error(err, "Could not create new ttl controller")
		return reconcile.Result{}, nil
	}

	mc.parentControllers[ctrlName] = ctrl
	go func() {
		mc.logger.Info("Starting TTLPolicy controller", "Name", ctrlName)
		if err := ctrl.Start(); err != nil {
			mc.logger.Error(err, "Could not start ttl controller")
		}
	}()

	return reconcile.Result{}, nil
}

func (mc *metacontroller) deleteController(ctrlName string, ttlPolicy *v1alpha1.TTLPolicy) (reconcile.Result, error) {
	if ctrl, ok := mc.parentControllers[ctrlName]; ok {
		mc.logger.Info("Deleting controller", "Name", ctrlName)
		ctrl.Stop()
		delete(mc.parentControllers, ctrlName)
	}

	ttlPolicy.SetFinalizers([]string{})
	if err := mc.client.Update(context.TODO(), ttlPolicy); err != nil {
		return reconcile.Result{}, fmt.Errorf("could not remove finalizer from TTLPolicy: %+v", err)
	}
	return reconcile.Result{}, nil
}
