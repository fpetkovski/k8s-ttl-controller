package main

import (
	"os"

	"github.com/fpetkovski/k8s-ttl-controller/pkg/apis/fpetkovski_io/v1alpha1"
	"github.com/fpetkovski/k8s-ttl-controller/pkg/ttl_controller"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func main() {
	log.SetLogger(zap.New())

	logger := log.Log.WithName("main")

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

	// Setup a new controller to reconcile ReplicaSets
	c, err := controller.New("ttl-metacontroller", mgr, controller.Options{
		Reconciler: ttl_controller.NewMetacontroller(mgr, log.Log),
	})
	if err != nil {
		logger.Error(err, "unable to create a composite controller")
		os.Exit(1)
	}

	// Watch ReplicaSets and enqueue ReplicaSet object key
	watchSource := &source.Kind{Type: &v1alpha1.TTLPolicy{}}
	if err := c.Watch(watchSource, &handler.EnqueueRequestForObject{}); err != nil {
		logger.Error(err, "unable to watch composite controllers")
		os.Exit(1)
	}

	logger.Info("Starting manager")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		logger.Error(err, "unable to run manager")
		os.Exit(1)
	}
}
