# k8s-ttl-operator

The k8s-ttl-operator project is a Kubernetes operator which enables defining Time to live (TTL) policies for Kubernetes resources. The behavior for a TTL policy is defined using the following user-defined parameters:
* `expirationFrom` - defines which field of the resource to use as a reference time point. If not defined, it defaults to the resource's creation time
* `ttlFrom` determines how long the resource should live after it expires

## Install

Install the CRD using
```bash
kubectl apply -f https://raw.githubusercontent.com/fpetkovski/k8s-ttl-operator/0.3.0/deploy/crds.yaml
```
and the operator using
```bash
kubectl apply -f https://raw.githubusercontent.com/fpetkovski/k8s-ttl-operator/0.3.0/deploy/operator.yaml
```

## Example usage

Expiring completed Kubernetes jobs can be done by creating a `TTLPolicy` with the following configuration
```yaml
apiVersion: fpetkovski.io/v1alpha1
kind: TTLPolicy
metadata:
  name: jobs-ttl-controller
spec:
  resource:
    apiVersion: batch/v1
    kind: Job
  expirationFrom: .status.completionTime 
  ttlFrom: .metadata.annotations.ttl
```

Please refer to the [examples](https://github.com/fpetkovski/k8s-ttl-operator/tree/main/examples) folder for more usage patterns.

## License
[MIT](https://choosealicense.com/licenses/mit/)