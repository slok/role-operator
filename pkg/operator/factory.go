package operator

import (
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"

	"github.com/slok/role-operator/pkg/log"
	"github.com/slok/role-operator/pkg/service/k8s"
)

const (
	operatorName = "role-operator"
)

// New returns a new operator ready to be run.
func New(cfg Config, k8ssvc k8s.Service, logger log.Logger) operator.Operator {
	logger = logger.With("operator", operatorName)

	handler := NewHandler(logger)
	crd := NewMultiRoleBindingCRD(k8ssvc)
	ctrl := controller.NewSequential(cfg.ResyncDuration, handler, crd, nil, logger)
	return operator.NewOperator(crd, ctrl, logger)
}
