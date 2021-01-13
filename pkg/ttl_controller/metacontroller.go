package ttl_controller

import (
	"context"
	"fmt"
	"fpetkovski/k8s-ttl-operator/pkg/apis/fpetkovski_io/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"math/rand"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strconv"
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

func NewMetacontroller(mgr controllerruntime.Manager) *metacontroller {
	return &metacontroller{
		client:            mgr.GetClient(),
		manager:           mgr,
		parentControllers: make(map[string]*ttlController),
		logger:            log.Log.WithName("metacontroller"),
	}
}

// Implement reconcile.Reconciler so the controller can reconcile objects
var _ reconcile.Reconciler = &metacontroller{}

func (mc *metacontroller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	cc := &v1alpha1.TTLController{}
	if err := mc.client.Get(context.TODO(), request.NamespacedName, cc); errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not get composite controller: %+v", err)
	}

	parentCtrlName := request.NamespacedName.String() + "-" + strconv.Itoa(rand.Int())
	if cc.GetDeletionTimestamp() != nil {
		mc.logger.Info("Removing controller", "Name", request.NamespacedName)
		return mc.deleteController(parentCtrlName, cc)
	}

	if !controllerutil.ContainsFinalizer(cc, finalizer) {
		controllerutil.AddFinalizer(cc, finalizer)
		if err := mc.client.Update(context.TODO(), cc); err != nil {
			return reconcile.Result{}, fmt.Errorf("could not set finalizer on composite controller: %+v", err)
		}
	}

	return mc.syncController(parentCtrlName, cc)
}

func (mc *metacontroller) syncController(ctrlName string, cc *v1alpha1.TTLController) (reconcile.Result, error) {
	if ctrl, found := mc.parentControllers[ctrlName]; found {
		mc.logger.Info("TTL controller already exists, recreating", "Name", ctrlName)
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
	ctrl, err := newTTLController(ctrlName, mc.manager, gvk, cc.Spec.TTLValueField, cc.Spec.ExpirationValueField)
	if err != nil {
		mc.logger.Error(err, "Could not create new ttl controller")
		return reconcile.Result{}, nil
	}

	mc.parentControllers[ctrlName] = ctrl
	go func() {
		if err := ctrl.Start(); err != nil {
			mc.logger.Error(err, "Could not start ttl controller")
		}
	}()

	return reconcile.Result{}, nil
}

func (mc *metacontroller) deleteController(ctrlName string, cc *v1alpha1.TTLController) (reconcile.Result, error) {
	if ctrl, ok := mc.parentControllers[ctrlName]; ok {
		mc.logger.Info("Deleting composite controller", "Name", ctrlName)
		ctrl.Stop()
		delete(mc.parentControllers, ctrlName)
	}

	cc.SetFinalizers([]string{})
	if err := mc.client.Update(context.TODO(), cc); err != nil {
		return reconcile.Result{}, fmt.Errorf("could not set finalizer on composite controller: %+v", err)
	}
	return reconcile.Result{}, nil
}
