package ttl_controller

import (
	"github.com/fpetkovski/k8s-ttl-controller/pkg/apis/fpetkovski_io/v1alpha1"
	"github.com/fpetkovski/k8s-ttl-controller/pkg/signals"
	"github.com/fpetkovski/k8s-ttl-controller/pkg/watch_predicates"

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
	ttlPolicy v1alpha1.TTLPolicySpec,
) (*ttlController, error) {
	logger := log.Log.WithName(name)
	groupVersion, err := schema.ParseGroupVersion(ttlPolicy.ResourceRule.APIVersion)
	if err != nil {
		return nil, err
	}

	gvk := groupVersion.WithKind(ttlPolicy.ResourceRule.Kind)
	objectMatcher := watch_predicates.And(
		watch_predicates.NamespacePredicate(ttlPolicy.ResourceRule.Namespace),
		watch_predicates.MatchLabelsPredicate(ttlPolicy.ResourceRule.MatchLabels),
	)
	r := newReconciler(
		mgr.GetClient(),
		gvk,
		ttlPolicy.TTLFrom,
		ttlPolicy.ExpirationFrom,
		objectMatcher,
		logger,
	)
	ctrl, err := controller.NewUnmanaged(name, mgr, controller.Options{
		Reconciler: r,
	})
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := ctrl.Watch(
		&source.Kind{Type: u},
		&handler.EnqueueRequestForObject{},
		objectMatcher,
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
