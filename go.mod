module github.com/fpetkovski/k8s-ttl-controller

go 1.14

require (
	github.com/go-logr/logr v0.4.0
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/stretchr/testify v1.4.0
	k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.6
	k8s.io/klog/v2 v2.8.0
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // indirect
	sigs.k8s.io/controller-runtime v0.6.4
)
