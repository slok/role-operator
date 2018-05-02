package k8s

import (
	"github.com/spotahome/kooper/client/crd"
	apiextensionscli "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	roleoperatorv1alpha1 "github.com/slok/role-operator/pkg/apis/roleoperator/v1alpha1"
	roleoperatorcli "github.com/slok/role-operator/pkg/client/k8s/clientset/versioned"
	"github.com/slok/role-operator/pkg/log"
)

// MultiRoleBinding knows how to manage the CRD MultiRoleBinding.
type MultiRoleBinding interface {
	// EnsureCRD will register the CRD.
	MultiRoleBindingEnsureCRD() error
	// List will list MultiRoleBinding CRs.
	MultiRoleBindingList(namespace string, opts metav1.ListOptions) (*roleoperatorv1alpha1.MultiRoleBindingList, error)
	// Watch will watch MultiRoleBinding CRs.
	MultiRoleBindingWatch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
}

type multiRoleBinding struct {
	rocli  roleoperatorcli.Interface
	crdcli crd.Interface
	logger log.Logger
}

// NewMultiRoleBinding returns a new multiRoleBinding instance.
func NewMultiRoleBinding(aeClient apiextensionscli.Interface, rocli roleoperatorcli.Interface, logger log.Logger) MultiRoleBinding {
	logger = logger.With("service", "k8s.multiRoleBinding")
	crdcli := crd.NewClient(aeClient, logger)

	return &multiRoleBinding{
		rocli:  rocli,
		crdcli: crdcli,
		logger: logger,
	}
}

// EnsureCRD satisfies MultiRoleBinding interface.
func (m *multiRoleBinding) MultiRoleBindingEnsureCRD() error {
	crdConf := crd.Conf{
		NamePlural: roleoperatorv1alpha1.MultiRoleBindingNamePlural,
		Kind:       roleoperatorv1alpha1.MultiRoleBindingKind,
		Group:      roleoperatorv1alpha1.SchemeGroupVersion.Group,
		Version:    roleoperatorv1alpha1.SchemeGroupVersion.Version,
		Scope:      roleoperatorv1alpha1.MultiRoleBindingScope,
	}
	return m.crdcli.EnsurePresent(crdConf)
}

// List satisfies MultiRoleBinding interface.
func (m *multiRoleBinding) MultiRoleBindingList(namespace string, opts metav1.ListOptions) (*roleoperatorv1alpha1.MultiRoleBindingList, error) {
	return m.rocli.RoleoperatorV1alpha1().MultiRoleBindings(namespace).List(opts)
}

// Watch satisfies MultiRoleBinding interface.
func (m *multiRoleBinding) MultiRoleBindingWatch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return m.rocli.RoleoperatorV1alpha1().MultiRoleBindings(namespace).Watch(opts)
}
