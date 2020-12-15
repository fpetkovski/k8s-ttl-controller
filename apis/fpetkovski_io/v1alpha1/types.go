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

// +groupName=fpetkovski.io
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ttlcontrollers,scope=Cluster,shortName=ttlctl
type TTLController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   TTlControllerSpec   `json:"spec"`
	Status TTLControllerStatus `json:"status,omitempty"`
}

type TTlControllerSpec struct {
	Resource             ResourceRule `json:"resource"`
	TTLValueField        string       `json:"ttlValueField"`
	ExpirationValueField *string      `json:"expirationValueField,omitempty"`
}

type ResourceRule struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type TTLControllerStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TTLControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TTLController `json:"items"`
}
