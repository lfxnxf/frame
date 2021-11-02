package inits

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpclient "github.com/lfxnxf/frame/logic/inits/http/client"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
)

func createFile(fileType string) *os.File {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d.%s", time.Now().UnixNano(), fileType))
	fh, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	return fh
}

func TestInit(t *testing.T) {
	f1 := createFile("toml")
	content := `[server]
	service_name="a.b.c"
	port = 17891
[log]
    level="debug"
    rotate="hour"
    logpath = "./logs"
	succ_stat_code=[1,2,3,4]

[data_log]
    path="logs/trans.log"
    rotate="day"

`
	err := ioutil.WriteFile(f1.Name(), []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cl := first()
	cl2 := second()

	configS := flag.String("config", f1.Name(), "Configuration file")
	flag.Parse()

	Init(ConfigPath(*configS))

	srv := HTTPServer()
	srv.ANY("/hello", func(c *httpserver.Context) {
		c.Response.WriteString("hello ,this is from HelloServer func")
	})
	go func() {
		srv.Run()
	}()

	time.Sleep(1 * time.Second)

	logging.DataLog("topic", "a", "b")
	fmt.Println("final config:", Default.config)
	fmt.Printf("acc log:%+v\n", DefaultKit.A())
	fmt.Printf("bus log:%+v\n", DefaultKit.B())
	fmt.Printf("err log:%+v\n", DefaultKit.E())
	fmt.Printf("balance log:%+v\n", logging.Log(logging.BalanceLoggerName))
	fmt.Printf("gen log:%+v\n", logging.Log(logging.GenLoggerName))
	fmt.Printf("slow log:%+v\n", logging.Log(logging.SlowLoggerName))
	fmt.Printf("crash log:%+v\n", logging.Log(logging.CrashLoggerName))
	fmt.Printf("default log:%+v\n", logging.Log(logging.DefaultLoggerName))

	req := httpclient.NewRequest().WithPath("/hello")
	rsp, err := cl.Call(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.String())
	time.Sleep(1 * time.Second)

	rsp2, err := cl2.Call(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp2.String())

	logging.Sync()
	time.Sleep(1 * time.Second)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello ,this is from HelloServer func ")
}

func first() httpclient.Client {
	d := New()
	sc := ServerClient{
		ServiceName:   "a.b.c",
		Ipport:        "127.0.0.1:17891",
		Balancetype:   "WeightRoundRobin",
		ProtoType:     "http",
		EndpointsFrom: "consul",
	}
	d.InjectServerClient(sc)
	d.Init()
	return d.HTTPClient("a.b.c")
}

func second() httpclient.Client {
	d := New()
	sc := ServerClient{
		ServiceName:   "x.y.z",
		Ipport:        "127.0.0.1:17891",
		Balancetype:   "WeightRoundRobin",
		ProtoType:     "http",
		EndpointsFrom: "consul",
	}
	d.InjectServerClient(sc)
	d.Init()

	return d.HTTPClient("x.y.z")
}

func TestInits_initWithMemory(t *testing.T) {
	data := `[server]
	service_name="user.base.user_sms_base"
	port = 10002

[log]
	level="debug"
	logpath="logs"
	rotate="hour"

[data_log]
  path="./logs/bigdata.log"
  rotate="hour"

[interval]
	table_check = 86400
	cosul_check = 60
	marketing_check = 5
	receipt_check = 86400

[remote_path]
    account="/account"
    channel="/channel"
    template="/template"
    strategy="/strategy"
    common="/config.toml"


[[server_client]]
  service_name="inf.secret.base"
  proto="http"
  endpoints="10.55.3.187:20206"
  endpoints_from="consul"
  balancetype="roundrobin"
  read_timeout=1000
  retry_times=1
  slow_time = 200`

	d := New()
	d.Init(
		ConfigMemory([]byte(data)),
		Name("inf.sms.base"),
		App("inke"),
		Namespace("loadtest"),
	)

	fmt.Printf("%+v", d.configInstance)
}

