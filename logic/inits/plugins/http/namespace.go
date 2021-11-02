package httpplugin

import (
	"github.com/lfxnxf/frame/logic/inits"
	http "github.com/lfxnxf/frame/logic/inits/http/server"
)

func Namespace(c *http.Context) {
	if c.Namespace != "" {
		c.Ctx = inits.WithAPPKey(c.Ctx, c.Namespace)
	}
	c.Next()
}
