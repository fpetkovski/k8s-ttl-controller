# k8s-ttl-operator

The k8s-ttl-operator project is a Kubernetes operator which enables time-to-live (TTL) behavior for Kubernetes resources. Once the TTL period for a resource expires, the resource will be deleted from Kubernetes. 

Deletion behavior is controlled using the following user-defined parameters:
* `expirationValueField` which indicates which field of the resource to use as expiration time
* `ttlValueField` which indicates how long the resource should live after it expires

## Install

Install the CRD using
```bash
kubectl apply -f https://raw.githubusercontent.com/fpetkovski/k8s-ttl-operator/main/deploy/crds.yaml
```
and the operator using
```bash
kubectl apply -f https://raw.githubusercontent.com/fpetkovski/k8s-ttl-operator/main/deploy/operator.yaml
```

## Example usage

Expiring completed Kubernetes jobs can be done by creating a `TTLController` with the following configuration
```yaml
apiVersion: fpetkovski.io/v1alpha1
kind: TTLController
metadata:
  name: jobs-ttl-controller
spec:
  resource:
    apiVersion: batch/v1
    kind: Job
  expirationValueField: .status.completionTime # Which field from the resource to use as expiration time
  ttlValueField: .metadata.annotations.ttl # Which field from the job resource to use as TTL
```

Please refer to the [examples](https://github.com/fpetkovski/k8s-ttl-operator/tree/main/examples) folder for more usage patterns.

## License
[MIT](https://choosealicense.com/licenses/mit/)