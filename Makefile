.PHONY: gen
gen:
	controller-gen crd paths=./pkg/apis/... output:stdout > deploy/crds.yaml
	controller-gen object paths=./pkg/apis/...