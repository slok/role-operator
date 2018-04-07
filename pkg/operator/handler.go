package operator

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/slok/role-operator/pkg/log"
)

// Handler will handle the operator CRD events.
type Handler struct {
	logger log.Logger
}

func NewHandler(logger log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// Add satisfies controller.Handler interface.
func (h *Handler) Add(obj runtime.Object) error {
	h.logger.Infof("Add obj %#v", obj)
	return nil
}

// Delete satisfies controller.Handler interface.
func (h *Handler) Delete(objKey string) error {
	h.logger.Infof("Dele obj %s", objKey)
	return nil
}
