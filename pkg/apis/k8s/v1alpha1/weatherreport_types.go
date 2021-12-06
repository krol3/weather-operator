package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WeatherReportSpec defines the desired state of WeatherReport
// +k8s:openapi-gen=true
type WeatherReportSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	City     string `json:"city"`
	Days     int    `json:"days"`
	Replicas int32  `json:"replicas"`
}

// WeatherReportStatus defines the observed state of WeatherReport
// +k8s:openapi-gen=true
type WeatherReportStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	State string `json:"state"`
	Pod   string `json:"pod"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WeatherReport is the Schema for the weatherreports API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type WeatherReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WeatherReportSpec   `json:"spec"`
	Status WeatherReportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WeatherReportList contains a list of WeatherReport
type WeatherReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []WeatherReport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WeatherReport{}, &WeatherReportList{})
}
