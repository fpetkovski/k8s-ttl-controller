package main

import (
	"k8s.io/klog/v2/klogr"
	"os"
	controllerruntime "sigs.k8s.io/controller-runtime"

	"github.com/fpetkovski/k8s-ttl-controller/pkg/apis/fpetkovski_io/v1alpha1"
	"github.com/fpetkovski/k8s-ttl-controller/pkg/ttl_controller"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func main() {
	logger := klogr.New()
	controllerruntime.SetLogger(logger)

	// Setup a manager
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		logger.Error(err, "unable to create a controller manager")
		os.Exit(1)
	}
	err = v1alpha1.AddToScheme(mgr.GetScheme())
	if err != nil {
		logger.Error(err, "could not add v1alpha1 to schema")
		os.Exit(1)
	}

	// Setup a new controller to reconcile TTLPolicies
	c, err := controller.New("ttlpolicies", mgr, controller.Options{
		Reconciler: ttl_controller.NewMetacontroller(mgr, logger),
	})
	if err != nil {
		logger.Error(err, "unable to create TTLPolicies controller")
		os.Exit(1)
	}

	// Watch TTLPolicies
	watchSource := &source.Kind{Type: &v1alpha1.TTLPolicy{}}
	if err := c.Watch(watchSource, &handler.EnqueueRequestForObject{}); err != nil {
		logger.Error(err, "unable to watch TTLPolicies")
		os.Exit(1)
	}

	logger.Info("Starting manager")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		logger.Error(err, "unable to start manager")
		os.Exit(1)
	}
}
