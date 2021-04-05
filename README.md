# k8s-ttl-controller

The k8s-ttl-controller project is a Kubernetes controller which provides a time to live (TTL) mechanism for Kubernetes resources. The TTL behavior for each resource is configured through a `TTLPolicy` using the following user-defined parameters:
* `expirationFrom` - the resource property used as reference point from which expiration is calculated. If not specified, it defaults to the resource's creation time
* `ttlFrom` the resource property used as the TTL value

## Installation

Install the CRD using:
```bash
kubectl apply -f https://raw.githubusercontent.com/fpetkovski/k8s-ttl-controller/0.6.0/deploy/crds.yaml
```
and the controller using:
```bash
kubectl apply -f https://raw.githubusercontent.com/fpetkovski/k8s-ttl-controller/0.6.0/deploy/controller.yaml
```

## Examples

Expiring Kubernetes jobs in the `Completed` state can be done by creating a `TTLPolicy` with the following configuration:
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

With this configuration, the TTL for jobs will be taken from the `ttl` annotation. A job whose `ttl` annotation is set to `15s` will be deleted 15 seconds after it completes.

Please refer to the [examples](https://github.com/fpetkovski/k8s-ttl-controller/tree/main/examples) folder for more usage patterns.

## API Reference

A detailed API reference can be found on the [API docs](https://github.com/fpetkovski/k8s-ttl-controller/tree/main/docs) page.

## License
[MIT](https://choosealicense.com/licenses/mit/)