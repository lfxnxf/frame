package inits

import (
	"strings"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits/breaker"
	"github.com/lfxnxf/frame/logic/inits/config"
	"github.com/lfxnxf/frame/logic/inits/config/encoder/toml"
	ns "github.com/lfxnxf/frame/logic/inits/internal/kit/namespace"
	"github.com/lfxnxf/frame/logic/inits/ratelimit"
)

func initConfigWatcher(remotePath string) {
	logging.GenLogf("on config init watcher path:%s", remotePath)
	go func() {
		w := config.WatchPrefix(remotePath)
		for {
			v := w.Next()
			decodeConfig(v)
		}
	}()
}

func decodeConfig(data map[string]string) {
	if len(data) == 0 {
		return
	}

	if val, ok := data[loadtestRemotePath]; ok {
		GlobalNamespace.Add(
			ns.LoadtestNamespace,
			ConfigMemory([]byte(val)),
			App(Default.App),
			Name(Default.Name),
			Version(Default.Version),
			Deps(Default.Deps),
		)
	}

	// /app/loadtest/config.toml加载新的inits
	// /app/yimi/config.toml
	// /config.toml
	for k, v := range data {
		if !strings.HasSuffix(k, "config.toml") {
			continue
		}

		namespace := ""
		if k != "/config.toml" {
			ss := strings.Split(k, "/")
			if len(ss) != 4 {
				continue
			}
			namespace = ss[2]
		}

		d := initsConfig{}
		if err := toml.NewEncoder().Decode([]byte(v), &d); err != nil {
			logging.GenLogf("on config watcher, decode error: %v", err)
			return
		}

		logging.GenLogf("on config watcher, reload namespace:%s", namespace)
		brkConfigs := getBreakerConfig(namespace, d)
		breaker.ReloadConfig(brkConfigs)
		limConfigs := getLimiterConfig(namespace, d)
		ratelimit.ReloadConfig(limConfigs)
	}
}
