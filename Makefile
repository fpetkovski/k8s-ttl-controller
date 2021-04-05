.PHONY: gen
gen:
	controller-gen crd paths=./pkg/apis/... output:stdout > deploy/crds.yaml
	controller-gen object paths=./pkg/apis/...

.PHONY: docs
docs:
	./hack/docs/gen-crd-api-reference-docs -template-dir hack/docs/templates -config hack/docs/config.json -api-dir "github.com/fpetkovski/k8s-ttl-controller/pkg/apis/" -out-file docs/README.md
