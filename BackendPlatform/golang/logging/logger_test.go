package logging

import (
	"fmt"
	_ "log"
	_ "net/http/pprof"
	"os"
	_ "os/signal"
	"runtime"
	"testing"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/rolling"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestOldLogger(t *testing.T) {

	l := New()
	l.SetRotateByHour()
	//	l.SetPrintLevel(false)
	l.SetHighlighting(false)
	l.Infof("hahahah")
	time.Sleep(time.Second * 2)
	l.Debugf("hasdhdhash %s", 123)

	log := l.Logger()
	log.Printf("this logger")
	time.Sleep(time.Second * 2)
	l.SetLogPrefix(fmt.Sprintf("%d|", os.Getpid()))
	l.SetOutputByName("./testold.log")
	l.Info("this logger 2222")
	l.SetFlags(0)
	l.Info("this logger 2222")
	l.Info("this logger 2222")
	//	os.Remove("./testold.log")
}

func TestLoggerWithFiled(t *testing.T) {
	l := With("test", "test value", "key", "1", "value", "1")
	l.Debugw("hahhh logw url", "test", 1234)
	l.Debugf("hahhh logf %s url %d", "test", 1234)

	Debugw("debugw test message", "url", "http://service.inke.cn/serviceinfo", "timeout", 3, "retry", 10)

}

func TestDataLogger(t *testing.T) {
	InitData("./bigdata/trans.log", rolling.DailyRolling)
	DataLog("topic_test", "url", "http://service.inke.cn/serviceinfo/info", "timeout", 3, "retry", 10)
	DataLog("topic_test", "url", "http://service.inke.cn/serviceinfo/info", "timeout", 3, "retry", 10, "info", map[string]interface{}{"key": "value", "key2": "value2"})
}

func TestDataLoggerWithKey(t *testing.T) {
	InitDataWithKey("./bigdata/trans.log", rolling.DailyRolling, "test_bigdata")
	DataLog("topic_test", "url", "http://service.inke.cn/serviceinfo/info", "timeout", 3, "retry", 10)
	DataLog("topic_test", "url", "http://service.inke.cn/serviceinfo/info", "timeout", 3, "retry", 10, "info", map[string]interface{}{"key": "value", "key2": "value2"})
}

func TestNewLogger(t *testing.T) {
	l := NewLogger(&Options{
		//		DisableColors: true,
		//		DisableLevel: true,
		Rolling: SECONDLY,
		//		TimesFormat: time.RFC3339Nano,
		TimesFormat: TIMESECOND,
	}, "test1.log", "test2.log")

	l.SetLogPrefix("log_prefix")
	//l.SetOutputByName("./test.log")
	//	l.SetRotateBySecond()
	//	l.Info("hahahah")
	l.Debugf("hasdhdhash %d", 123)

	for i := 0; i < 700; i++ {
		Log("test2").Infof("hahahah %d", i)
	}
	//os.Remove("./test.log*")
}

func BenchmarkDebugLogParallel2(b *testing.B) {
	//fileobj, _ := os.OpenFile("test3.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	SetOutputByName("test3.log")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//Debugf("test log debug %d", 10234)
			Info("test log debug")
			Sync()
			//Debugf("test log info sadadadad")
			//_defaultLogger.Sync()
			//for _, f := range _defaultLogger.rollingFiles {
			//	f.(*rollingFile).Sync()
			//}
			//fileobj.WriteString(fmt.Sprintf("test log debug %d", 1234))
		}
	})
}

func BenchmarkDebugLogParallelZap(b *testing.B) {
	fileobj, _ := os.OpenFile("test3.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	sugar := zap.New(zapcore.NewCore(
		enc,
		zapcore.Lock(fileobj),
		zap.DebugLevel,
	)).Sugar()

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//Debugf("test log debug %d", 10234)
			sugar.Debugf("test log debug %d", 1234)
			//fileobj.WriteString(fmt.Sprintf("test log debug %d", 1234))
		}
	})
}

func BenchmarkDebugLogParallel(b *testing.B) {
	//fileobj, _ := os.OpenFile("test3.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	SetOutputByName("test3.log")
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Debugf("test log debug %d", 10234)
			//Info("test log debug")
			//fileobj.WriteString(fmt.Sprintf("test log debug %d", 1234))
		}
	})
}

func BenchmarkDebugLog(b *testing.B) {
	SetOutputByName("test.log")
	for i := 0; i < b.N; i++ {
		//Debugf("test log debug")
		Infof("test log debug")
	}
}

func TestDefaultLog(t *testing.T) {
	//InitError("log")
	log := NewLogger(&Options{})
	log.Errorf("%s", "this is error")
}

func setUpLogger() {
	cc := CommonLogConfig{
		Pathprefix:      "logs/",
		Rotate:          "day",
		GenLogLevel:     "info",
		BalanceLogLevel: "info",
	}
	SetOutputPath("logs/")
	InitCommonLog(cc)

}
func BenchmarkCommonErrorLog(b *testing.B) {
	setUpLogger()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Errorf("test log debug")
		}
	})
}
func BenchmarkCommonDebugLog(b *testing.B) {
	setUpLogger()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Debugf("test log debug")
		}
	})
}

func BenchmarkRuntimeStack(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bt := make([]byte, 4096)
			runtime.Stack(bt, false)
		}
	})
}

// func BenchmarkDebugLog(b *testing.B) {
// 	SetOutputByName("test.log")
// 	for i := 0; i < b.N; i++ {
// 		Debugf("test log debug")
// 	}
// }

// func BenchmarkDebugLogParallel(b *testing.B) {
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			Debugf("test log debug %d", 10234)
// 		}
// 	})
// }
