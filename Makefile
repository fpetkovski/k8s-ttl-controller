.PHONY: gen
gen:
	controller-gen crd paths=./pkg/apis/... output:stdout > deploy/crds.yaml
	controller-gen object paths=./pkg/apis/...

.PHONY: docs
docs:
	go mod vendor
	./scripts/docs/gen-crd-api-reference-docs -template-dir scripts/docs/templates -config scripts/docs/config.json -api-dir "github.com/fpetkovski/k8s-ttl-controller/pkg/apis/" -out-file docs/README.md
	rm -rf vendor

.PHONY: fmt
fmt:
	goimports -w -l .

