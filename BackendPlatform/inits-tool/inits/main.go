package main

import (
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/cmd"
)

func main() {
	go func() {
		_ = cmd.ToolStat()
	}()

	cmd.Execute()
}
