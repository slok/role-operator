package operator

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	roleoperatorv1alpha1 "github.com/slok/role-operator/pkg/apis/roleoperator/v1alpha1"
	"github.com/slok/role-operator/pkg/service/k8s"
)

// MultiRoleBindingCRD knows how to register and retrieve MultiRoleBinding CRs.
type MultiRoleBindingCRD struct {
	k8ssvc k8s.Service
}

// NewMultiRoleBindingCRD returns a new MultiRoleBindingCRD.
func NewMultiRoleBindingCRD(k8ssvc k8s.Service) *MultiRoleBindingCRD {
	return &MultiRoleBindingCRD{
		k8ssvc: k8ssvc,
	}
}

// Initialize satisfies crd.Interface.
func (m *MultiRoleBindingCRD) Initialize() error {
	return m.k8ssvc.MultiRoleBindingEnsureCRD()
}

// GetListerWatcher satisfies crd.Interface.
func (m *MultiRoleBindingCRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return m.k8ssvc.MultiRoleBindingList("", options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return m.k8ssvc.MultiRoleBindingWatch("", options)
		},
	}
}

// GetObject satisfies crd.Interface.
func (m *MultiRoleBindingCRD) GetObject() runtime.Object {
	return &roleoperatorv1alpha1.MultiRoleBinding{}
}
