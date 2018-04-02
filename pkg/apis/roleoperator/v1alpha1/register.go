package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/slok/role-operator/pkg/apis/roleoperator"
)

const (
	version = "v1alpha1"
)

// Workspace constants
const (
	PermissionKind       = "Permission"
	PermissionName       = "permission"
	PermissionNamePlural = "permissions"
	PermissionScope      = apiextensionsv1beta1.ClusterScoped
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: roleoperator.GroupName, Version: version}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Permission{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
