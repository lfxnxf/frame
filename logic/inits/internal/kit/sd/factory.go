package sd

import (
	"github.com/lfxnxf/frame/logic/inits/internal/core"
)

type Factory interface {
	Factory(host string) (core.Plugin, error)
}
