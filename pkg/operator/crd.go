package operator

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

// MultiRoleBindingCRD knows how to register and retrieve MultiRoleBinding CRs.
type MultiRoleBindingCRD struct{}

func NewMultiRoleBindingCRD() *MultiRoleBindingCRD {
	return &MultiRoleBindingCRD{}
}

// Initialize satisfies crd.Interface.
func (m *MultiRoleBindingCRD) Initialize() error {
	return nil
}

// GetListerWatcher satisfies crd.Interface.
func (m *MultiRoleBindingCRD) GetListerWatcher() cache.ListerWatcher {
	return nil
}

// GetObject satisfies crd.Interface.
func (m *MultiRoleBindingCRD) GetObject() runtime.Object {
	return nil
}
