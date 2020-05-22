package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReverseWordsAppSpec defines the desired state of ReverseWordsApp
// +k8s:openapi-gen=true
type ReverseWordsAppSpec struct {
	Replicas int32	`json:"replicas"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// ReverseWordsAppStatus defines the observed state of ReverseWordsApp
// +k8s:openapi-gen=true
type ReverseWordsAppStatus struct {
	AppPods	[]string `json:"appPods"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ReverseWordsApp is the Schema for the reversewordsapps API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=reversewordsapps,scope=Namespaced
type ReverseWordsApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReverseWordsAppSpec   `json:"spec,omitempty"`
	Status ReverseWordsAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ReverseWordsAppList contains a list of ReverseWordsApp
type ReverseWordsAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReverseWordsApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReverseWordsApp{}, &ReverseWordsAppList{})
}
