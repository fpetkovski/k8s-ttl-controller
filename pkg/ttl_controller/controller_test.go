package ttl_controller_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/fpetkovski/k8s-ttl-controller/pkg/apis/fpetkovski_io/v1alpha1"
	"github.com/fpetkovski/k8s-ttl-controller/pkg/ttl_controller"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	defaultScheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2/klogr"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func TestController(t *testing.T) {
	setupEnv()

	env := envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			Paths: []string{"../../deploy"},
		},
	}
	cfg, err := env.Start()
	assert.NoError(t, err)
	defer env.Stop()

	mgr := setupManager(t, cfg)
	go func() {
		stopChan := make(chan struct{})
		mgr.Start(stopChan)
	}()

	pod := v1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
			Annotations: map[string]string{
				"ttl": "3s",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx",
					Image: "nginx:latest",
				},
			},
		},
	}

	k8s := mgr.GetClient()
	err = k8s.Create(context.TODO(), &pod)
	assert.NoError(t, err)

	ttlPolicy := v1alpha1.TTLPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "fpetkovski.io/v1alpha1",
			Kind:       "TTLPolicy",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-policy",
			Namespace: "default",
		},
		Spec: v1alpha1.TTLPolicySpec{
			ResourceRule: v1alpha1.ResourceRule{
				APIVersion: "v1",
				Kind:       "Pod",
			},
			TTLFrom: ".metadata.annotations.ttl",
		},
	}
	err = k8s.Create(context.TODO(), &ttlPolicy)
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		pod := v1.Pod{}
		key := types.NamespacedName{
			Namespace: "default",
			Name:      "nginx",
		}
		err := k8s.Get(context.TODO(), key, &pod)
		return errors.IsNotFound(err)
	}, 20*time.Second, time.Second)
}

func setupEnv() {
	if err := os.Setenv("TEST_ASSET_KUBE_APISERVER", "../../test/bin/kube-apiserver"); err != nil {
		log.Fatal(err.Error())
	}
	if err := os.Setenv("TEST_ASSET_ETCD", "../../test/bin/etcd"); err != nil {
		log.Fatal(err.Error())
	}
}

func setupManager(t *testing.T, cfg *rest.Config) manager.Manager {
	scheme := runtime.NewScheme()
	err := defaultScheme.AddToScheme(scheme)
	assert.NoError(t, err)
	err = v1alpha1.AddToScheme(scheme)
	assert.NoError(t, err)
	err = apiextensionsv1.AddToScheme(scheme)
	assert.NoError(t, err)

	mgr, err := manager.New(cfg, manager.Options{
		Scheme: scheme,
	})
	assert.NoError(t, err)

	logger := klogr.New()
	controllerruntime.SetLogger(logger)
	c, err := controller.New("ttlpolicies", mgr, controller.Options{
		Reconciler: ttl_controller.NewMetacontroller(mgr, logger),
	})
	assert.NoError(t, err)

	// Watch TTLPolicies
	watchSource := &source.Kind{Type: &v1alpha1.TTLPolicy{}}
	err = c.Watch(watchSource, &handler.EnqueueRequestForObject{})
	assert.NoError(t, err)

	return mgr
}
