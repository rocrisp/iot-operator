/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CameraSpec defines the desired state of Camera
type CameraSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Size       int32  `json:"size,omitempty"`
	Cameratype string `json:"cameratype"`
	Location   string `json:"location"`
	Setting    string `json:"setting,omitempty"`
}

// CameraStatus defines the observed state of Camera
type CameraStatus struct {
	// Represents the observations of a Camera's current state.
	// Camera.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// Camera.status.conditions.status are one of True, False, Unknown.
	// Conditions store the status conditions of the Camera instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Camera is the Schema for the cameras API
type Camera struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CameraSpec   `json:"spec,omitempty"`
	Status CameraStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CameraList contains a list of Camera
type CameraList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Camera `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Camera{}, &CameraList{})
}
