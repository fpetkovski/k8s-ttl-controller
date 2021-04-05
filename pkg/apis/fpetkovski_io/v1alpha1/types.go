/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ttlpolicies,scope=Cluster,shortName=ttlp
// TTLPolicy is the object through which time to live behavior is configured for a Kubernetes resource.
type TTLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// TTLPolicySpec is the spec of the TTLPolicy
	Spec TTLPolicySpec `json:"spec"`
	// TTLPolicySpec is the status of the TTLPolicy
	Status TTLPolicyStatus `json:"status,omitempty"`
}

type TTLPolicySpec struct {
	// ResourceRule defines the resources to which the TTLPolicy should be applied
	ResourceRule ResourceRule `json:"resource"`
	// TTLFrom is the resources' property which contains the TTL value for the specific resource. <br/>
	// Examples: <br />
	// - 15s <br />
	// - 1m <br />
	// - 1h30m <br />
	TTLFrom string `json:"ttlFrom"`
	// ExpirationFrom is the resources' property which contains the time from which TTL is calculated.
	// Examples include `.metadata.creationTimestamp` or `.status.startTime`.
	// The time should be specified in in `RFC3339` format.
	// +optional
	ExpirationFrom *string `json:"expirationFrom,omitempty"`
}

// ResourceRule defines the resources to which the TTLPolicy should be applied
type ResourceRule struct {
	// APIVersion is the full API version of the kubernetes resources. <br />
	// Examples: <br />
	// - v1 <br />
	// - apps/v1 <br />
	APIVersion string `json:"apiVersion"`
	// Kind is the resources' Kind.
	// Examples: <br />
	// - Deployment <br />
	// - Ingress <br />
	Kind string `json:"kind"`
	// Namespace is the namespace in which the resources are created
	// +optional
	Namespace *string `json:"namespace,omitempty"`
	// MatchLabels is the label set which the resources should match
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

type TTLPolicyStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TTLPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TTLPolicy `json:"items"`
}
