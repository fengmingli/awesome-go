/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HxlvmSpec defines the desired state of Hxlvm
type HxlvmSpec struct {
	Name   string `json:"name"`
	Size   string `json:"size"`
	VgName string `json:"vgName"`
}

// HxlvmStatus defines the observed state of Hxlvm
type HxlvmStatus struct {
	PvItems  map[string]Item `json:"pvItems,omitempty"`
	PvcItems map[string]Item `json:"pvcItems,omitempty"`
}
type Item struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	SizeDesc string `json:"sizeDesc"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Hxlvm is the Schema for the hxlvms API
type Hxlvm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HxlvmSpec   `json:"spec,omitempty"`
	Status HxlvmStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HxlvmList contains a list of Hxlvm
type HxlvmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hxlvm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hxlvm{}, &HxlvmList{})
}
