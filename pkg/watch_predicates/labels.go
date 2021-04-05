package watch_predicates

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type matchLabelsPredicate struct {
	matchLabels map[string]string
}

func MatchLabelsPredicate(labels map[string]string) *matchLabelsPredicate {
	return &matchLabelsPredicate{
		matchLabels: labels,
	}
}

func (n matchLabelsPredicate) Create(event event.CreateEvent) bool {
	return n.Matches(event.Meta)
}

func (n matchLabelsPredicate) Delete(event event.DeleteEvent) bool {
	return n.Matches(event.Meta)
}

func (n matchLabelsPredicate) Update(event event.UpdateEvent) bool {
	return n.Matches(event.MetaNew)
}

func (n matchLabelsPredicate) Generic(event event.GenericEvent) bool {
	return n.Matches(event.Meta)
}

func (n matchLabelsPredicate) Matches(object metav1.Object) bool {
	objectLabels := object.GetLabels()
	for k, v := range n.matchLabels {
		if objectLabels[k] != v {
			return false
		}
	}
	return true
}
