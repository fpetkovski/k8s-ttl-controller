package watch_predicates

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type ObjectMatcher interface {
	predicate.Predicate
	Matches(object metav1.Object) bool
}
