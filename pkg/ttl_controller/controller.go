package ttl_controller

import (
	"fpetkovski/k8s-ttl-operator/pkg/signals"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type ttlController struct {
	controller    controller.Controller
	stopChan      chan struct{}
	logger        logr.Logger
	ttlValueField string
}

func newTTLController(
	name string,
	mgr controllerruntime.Manager,
	gvk schema.GroupVersionKind,
	ttlValueField string,
	expirationValueField *string,
) (*ttlController, error) {
	logger := log.Log.WithName(name)
	ctrl, err := controller.NewUnmanaged(name, mgr, controller.Options{
		Reconciler: newReconciler(mgr.GetClient(), gvk, ttlValueField, expirationValueField, logger),
	})
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := ctrl.Watch(
		&source.Kind{Type: u},
		&handler.EnqueueRequestForObject{},
	); err != nil {
		return nil, err
	}

	return &ttlController{
		controller: ctrl,
		stopChan:   signals.SetupSignalHandler(),
		logger:     logger,
	}, nil
}

func (r *ttlController) Start() error {
	return r.controller.Start(r.stopChan)
}

func (r *ttlController) Stop() {
	r.stopChan <- struct{}{}
	return
}
