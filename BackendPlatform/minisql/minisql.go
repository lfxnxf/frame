package minisql

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	binPath         = "/usr/local/bin"
	file            = fmt.Sprintf("tidb-server_%s", runtime.GOOS)
	tarFile         = fmt.Sprintf("%s.tar.gz", file)
	filePath        = fmt.Sprintf("%s/%s", binPath, file)
	anotherFilePath = fmt.Sprintf("/build/dependency/bin/%s", file)
	tarFilePath     = fmt.Sprintf("%s/%s", binPath, tarFile)
	url             = fmt.Sprintf("http://m4a.inke.cn/%s", tarFile)
)

type MiniSQL struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
	e        *Cmd
}

// Run create *MiniSQL and start,then register mock-mysql driver
func Run() *MiniSQL {
	s := New()
	s.Start()
	return s
}

func New() *MiniSQL {
	return &MiniSQL{
		User:     "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     getPort(),
		DB:       "test",
	}
}

// Start start a mysql instance
func (m *MiniSQL) Start() (err error) {
	var (
		dataPath = fmt.Sprintf("/tmp/tidb-%d", time.Now().UnixNano())
		logPath  = fmt.Sprintf("%s/log.log", dataPath)
	)
	err = exists(dataPath)
	if err != nil {
		return
	}
	port := fmt.Sprintf("%d", m.Port)
	m.e, err = Start(getBin(), []string{"-P", port, "-path", dataPath, "-log-file", logPath})
	// wait tidb bootstrap
	time.Sleep(time.Second * 2)
	return
}

// Stop stop mysql instance
func (m *MiniSQL) Stop() error {
	return m.e.Stop()
}
func exists(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		err = os.RemoveAll(path)
		return
	}
	return nil
}
func getPort() (port int) {
	l, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		return
	}
	port = l.Addr().(*net.TCPAddr).Port
	err = l.Close()
	if err != nil {
		return
	}
	return
}

func getBin() string {
	_, err := os.Stat(filePath)
	_, existsErr := os.Stat(anotherFilePath)

	if err == nil || existsErr == nil || isInstall(file) {
		return file
	}
	err = download()
	if err != nil {
		fmt.Printf("download %v failed \n", url)
	}
	return file
}
func isInstall(exe string) bool {
	_, err := exec.LookPath(exe)
	if err != nil {
		return false
	}
	return true
}
func download() (err error) {
	var resp *http.Response
	var buf []byte
	client := http.Client{Timeout: time.Second * 300}
	_, err = os.Stat(tarFilePath)
	if err != nil {
		fmt.Printf("downloading mysql server,waiting 60s ...\n")
		resp, err = client.Get(url)
		if err != nil {
			return
		}
		buf, err = ioutil.ReadAll(resp.Body)
		err = ioutil.WriteFile(tarFilePath, buf, 0600)
		if err != nil {
			return
		}
	}
	_, err = Exec("tar", []string{"xzvf", tarFilePath, "-C", binPath})
	if err != nil {
		return
	}
	_, err = Exec("chmod", []string{"+x", filePath})
	return
}
