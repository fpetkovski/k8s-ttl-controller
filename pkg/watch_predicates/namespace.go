package watch_predicates

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type namespacePredicate struct {
	namespace *string
}

func NamespacePredicate(namespace *string) *namespacePredicate {
	return &namespacePredicate{namespace: namespace}
}

func (w namespacePredicate) Create(event event.CreateEvent) bool {
	return w.Matches(event.Meta)
}

func (w namespacePredicate) Delete(event event.DeleteEvent) bool {
	return w.Matches(event.Meta)
}

func (w namespacePredicate) Update(event event.UpdateEvent) bool {
	return w.Matches(event.MetaNew)
}

func (w namespacePredicate) Generic(event event.GenericEvent) bool {
	return w.Matches(event.Meta)
}

func (w namespacePredicate) Matches(object metav1.Object) bool {
	if w.namespace == nil {
		return true
	}

	return object.GetNamespace() == *w.namespace
}
