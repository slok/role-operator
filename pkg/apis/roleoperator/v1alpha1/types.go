package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=multirolebindings

// MultiRoleBinding is a specification of a rolebinding that will be multiplied and binded
// to the specified serviceAccount across multiple namespaces dynamically based on a selector.
type MultiRoleBinding struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the MultiRoleBinding spec.
	Spec MultiRoleBindingSpec `json:"spec,omitempty"`
}

// MultiRoleBindingSpec contains the specification for the MultiRoleBinding.
type MultiRoleBindingSpec struct {
	// RoleBindingRef is the specification of the desired behavior of the MultiRoleBinding.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	RoleBindingRef RoleBindingRef `json:"roleBindingRef,omitempty"`

	// NamespaceSelector is a label query over namespaces that should match in order to
	// grant access with the selected MultiRoleBinding from the referenced roleBinding.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	NamespaceSelector *metav1.LabelSelector `json:"selector"`
}

// RoleBindingRef contains the information to reference a rolebinding in namespace.
type RoleBindingRef struct {
	// Name of the rolebinding being referenced.
	Name string `json:"name"`
	// Namespace of the referenced role binding.
	Namespace string `json:"namespace,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MultiRoleBindingList is a collection of MultiRoleBindings.
type MultiRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultiRoleBinding `json:"items"`
}
