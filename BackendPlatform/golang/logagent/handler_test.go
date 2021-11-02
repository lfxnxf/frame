package logagent

import (
	"testing"
)

var agentclient *LogAgent

func TestTimerLogMsg(t *testing.T) {

	agentclient = InitAgent()
	remoteService := "live.liveinfo" // 调用远端服务
	metric := "getLiveInfo"          // 监控的信息,
	var duration int64 = 1000        // 耗时
	code := 0                        // 错误码
	agentclient.SendTimerLogMsg(metric, remoteService, metric, duration, code)
}

func TestTimerLog(t *testing.T) {
	agentclient = InitAgent()
	remoteService := "live.liveinfo" // 调用远端服务
	metric := "getLiveInfo"          // 监控的信息,
	var duration int64 = 1000        // 耗时
	code := 0                        // 错误码
	lat := NewLogAgentTagOptions()
	lat.SetTag("key", "value")
	agentclient.SendTimerLog(metric, remoteService, metric, duration, code, lat)
}

func TestSnapshot(t *testing.T) {
	agentclient = InitAgent()

	metric := "getLiveInfo"         // 监控的信息,
	var snapshotValue float64 = 100 //  快照数量
	lat := NewLogAgentTagOptions()
	lat.SetMemCache("liveidcount") // 设置当前是统计内存数量的tag 模式

	// lat.SetTag("key", "value") //自定义tag类型
	agentclient.SendSnapshotLog(metric, snapshotValue, lat)
}
