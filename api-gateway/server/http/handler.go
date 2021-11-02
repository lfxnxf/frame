package http

import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
)

func ping(c *httpserver.Context) {
	if err := svc.Ping(c.Ctx); err != nil {
		c.JSONAbort(nil, err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(okMsg, nil)
}

