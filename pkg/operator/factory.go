package operator

import (
	"github.com/slok/role-operator/pkg/log"

	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
)

const (
	operatorName = "role-operator"
)

// New returns a new operator ready to be run.
func New(cfg Config, logger log.Logger) operator.Operator {
	logger = logger.With("operator", operatorName)

	handler := NewHandler(logger)
	crd := NewMultiRoleBindingCRD()
	ctrl := controller.NewSequential(cfg.ResyncDuration, handler, crd, nil, logger)
	return operator.NewOperator(crd, ctrl, logger)
}
