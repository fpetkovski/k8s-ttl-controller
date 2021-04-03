package ttl_controller

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

// reconciler reconciles the parents from a composite controller
type reconciler struct {
	client               client.Client
	gvk                  schema.GroupVersionKind
	ttlValueField        string
	expirationValueField *string
	logger               logr.Logger
}

// Implement reconcile.Reconciler so the controller can reconcile objects
var _ reconcile.Reconciler = &reconciler{}

func newReconciler(
	client client.Client,
	gvk schema.GroupVersionKind,
	ttlValueField string,
	expirationValueField *string,
	logger logr.Logger,
) *reconciler {
	return &reconciler{
		client:               client,
		gvk:                  gvk,
		ttlValueField:        ttlValueField,
		expirationValueField: expirationValueField,
		logger:               logger,
	}
}

func (r *reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	resource := unstructured.Unstructured{}
	resource.SetGroupVersionKind(r.gvk)
	if err := r.client.Get(context.TODO(), request.NamespacedName, &resource); errors.IsNotFound(err) {
		r.logger.V(5).Info(
			"Could not find object, skipping.",
			"ApiVersion", r.gvk.GroupVersion(),
			"Kind", resource.GetKind(),
			"Name", request.Name)
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not get resource: %+v", err)
	}

	// Skip already deleted resources
	if resource.GetDeletionTimestamp() != nil {
		return reconcile.Result{}, nil
	}

	expirationTime := resource.GetCreationTimestamp().Time
	if r.expirationValueField != nil {
		t, err := GetExpirationValue(resource, *r.expirationValueField)
		if err != nil {
			r.logger.V(5).Info(
				fmt.Sprintf("Expiration value is not a valid time: %s", err.Error()),
				"Kind", resource.GetKind(),
				"Name", resource.GetName(),
				"ExpirationFrom", *r.expirationValueField)
			return reconcile.Result{}, nil
		} else {
			expirationTime = t
		}
	}

	ttl, err := GetTTLValue(resource, r.ttlValueField)
	if err != nil {
		r.logger.V(5).Info(
			fmt.Sprintf("Expiration value is not a valid duration: %s", err.Error()),
			"Error processing object",
			"Kind", resource.GetKind(),
			"Name", resource.GetName())
		return reconcile.Result{}, nil
	}

	if IsExpired(ttl, expirationTime) {
		return r.delete(resource)
	} else {
		return r.requeue(ttl, expirationTime, resource)
	}
}

func (r *reconciler) delete(resource unstructured.Unstructured) (reconcile.Result, error) {
	r.logger.Info("Object expired", "Kind", resource.GetKind(), "Name", resource.GetName())
	backgroundDeletion := client.PropagationPolicy(v1.DeletePropagationBackground)
	err := r.client.Delete(context.TODO(), &resource, backgroundDeletion)
	if err != nil {
		return reconcile.Result{
			RequeueAfter: 30 * time.Second,
		}, err
	}

	return reconcile.Result{}, nil
}

func (r *reconciler) requeue(ttl time.Duration, createdAt time.Time, resource unstructured.Unstructured) (reconcile.Result, error) {
	requeueAfter := createdAt.Add(ttl).Sub(time.Now())
	message := fmt.Sprintf("Scheduling deletion in %d seconds", int64(requeueAfter.Seconds()))
	r.logger.Info(message, "Kind", resource.GetKind(), "Name", resource.GetName())

	return reconcile.Result{
		RequeueAfter: requeueAfter,
	}, nil
}
