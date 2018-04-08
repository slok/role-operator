package k8s

import (
	roleoperatorcli "github.com/slok/role-operator/pkg/client/k8s/clientset/versioned"
	"github.com/slok/role-operator/pkg/log"
	apiextensionscli "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

// Service is the main way of managing k8s resources.
type Service interface {
	MultiRoleBinding
}

type service struct {
	MultiRoleBinding
}

// New returns a new k8s service.
func New(aeClient apiextensionscli.Interface, rocli roleoperatorcli.Interface, logger log.Logger) Service {
	return service{
		MultiRoleBinding: NewMultiRoleBinding(aeClient, rocli, logger),
	}
}
