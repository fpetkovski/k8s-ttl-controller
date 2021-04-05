package watch_predicates

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type and struct {
	matchers []ObjectMatcher
}

func (a and) Create(event event.CreateEvent) bool {
	for _, m := range a.matchers {
		if !m.Create(event) {
			return false
		}
	}
	return true
}

func (a and) Delete(event event.DeleteEvent) bool {
	for _, m := range a.matchers {
		if !m.Delete(event) {
			return false
		}
	}
	return true
}

func (a and) Update(event event.UpdateEvent) bool {
	for _, m := range a.matchers {
		if !m.Update(event) {
			return false
		}
	}
	return true
}

func (a and) Generic(event event.GenericEvent) bool {
	for _, m := range a.matchers {
		if !m.Generic(event) {
			return false
		}
	}
	return true
}

func (a and) Matches(object metav1.Object) bool {
	for _, m := range a.matchers {
		if !m.Matches(object) {
			return false
		}
	}

	return true
}

func And(matches ...ObjectMatcher) *and {
	return &and{
		matchers: matches,
	}
}
