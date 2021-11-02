package testcli

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
)

var (
	ymlPath     string
	serviceName string
	services    = make(map[string]*Container)
)

func Setup(confPath string) (testConfPath string) {
	var err error
	defer func() {
		if err != nil {
			logging.Errorf("setup %s error: %s", confPath, err)
			panic(err)
		}
	}()
	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		return
	}
	var conf Config
	err = toml.Unmarshal(buf, &conf)
	if err != nil {
		return
	}
	var dbName string
	if len(conf.Database) != 0 {
		var u *mysql.Config
		u, err = mysql.ParseDSN(conf.Database[0].Master)
		if err != nil {
			return
		}
		dbName = u.DBName
	}
	var kafkaTopic = []string{"test_topic:1:1"}
	for _, v := range conf.KafkaProducerClient {
		t := strings.Split(v.Topic, ",")
		for _, vv := range t {
			kafkaTopic = append(kafkaTopic, fmt.Sprintf("%s:1:1", vv))
		}

	}
	topic := strings.Join(kafkaTopic, ",")
	ip, err := HostIP()
	if err != nil || ip == "" {
		ip = "127.0.0.1"
	}
	yamlBuf := fmt.Sprintf(yml, dbName, ip, topic)
	yamlPath := path.Dir(confPath) + "/.docker-compose.yml"
	err = ioutil.WriteFile(yamlPath, []byte(yamlBuf), 0600)
	logging.Debugf("generate docker-compose.yml %s", yamlPath)
	if err != nil {
		return
	}
	re, _ := regexp.Compile("[.|_|-]")
	serviceName = re.ReplaceAllString(conf.Server.ServiceName, "")
	err = setupYAML(yamlPath)
	if err != nil {
		return
	}
	_, err = getServices()
	if err != nil {
		return
	}
	var result map[string]interface{}
	err = toml.Unmarshal(buf, &result)
	if err != nil {
		return
	}
	for k := range result {
		data, ok := result[k].([]map[string]interface{})
		if !ok {
			continue
		}
		switch k {
		case "database":
			for i, val := range data {
				dsn, _ := val["master"].(string)
				u, err := mysql.ParseDSN(dsn)
				if err != nil {
					continue
				}
				u.User = "root"
				u.Passwd = "root"
				u.Addr = fmt.Sprintf("127.0.0.1:%s", getPort("db"))
				u.DBName = dbName
				newDsn := u.FormatDSN()
				data[i]["master"] = newDsn
				data[i]["slaves"] = []string{newDsn}

			}
		case "redis":
			for i := range data {
				data[i]["addr"] = fmt.Sprintf("127.0.0.1:%s", getPort("redis"))
				data[i]["password"] = ""
			}
		case "kafka_producer_client":
			for i := range data {
				data[i]["kafka_broken"] = fmt.Sprintf("127.0.0.1:%s", getPort("kafka"))
			}
		}
	}
	buffer := &bytes.Buffer{}
	e := toml.NewEncoder(buffer)
	err = e.Encode(result)
	if err != nil {
		return
	}
	testConfPath = path.Dir(confPath) + "/.config.toml"
	ioutil.WriteFile(testConfPath, buffer.Bytes(), 0600)
	logging.Debugf("rewrite config.toml %s", testConfPath)
	time.Sleep(time.Millisecond * 500)
	checkServices()
	return
}

func runCompose(args ...string) (output []byte, err error) {
	if _, err = os.Stat(ymlPath); os.IsNotExist(err) {
		logging.Errorf("os.Stat(%s) composer yaml is not exist!", ymlPath)
		return
	}
	if ymlPath, err = filepath.Abs(ymlPath); err != nil {
		logging.Errorf("filepath.Abs(%s) error(%v)", ymlPath, err)
		return
	}
	args = append([]string{"-f", ymlPath, "-p", serviceName}, args...)
	logging.Debugf("exec command: docker-compose %s", strings.Join(args, " "))
	if output, err = exec.Command("docker-compose", args...).CombinedOutput(); err != nil {
		logging.Errorf("exec.Command(docker-compose) args(%v) stdout(%s) error(%v)", args, string(output), err)
		return
	}
	return
}

// Setup setup UT related environment dependence for everything.
func setupYAML(path string) (err error) {
	ymlPath = path
	// 先检查是否已经有network

	logging.Infof("setup container... ")
	if _, err = runCompose("up", "-d"); err != nil {
		return
	}
	defer func() {
		if err != nil {
			Teardown()
		}
	}()
	return
}

// Teardown unsetup all environment dependence.
func Teardown() (err error) {
	logging.Infof("teardown container... ")
	_, err = runCompose("down")
	return
}

func getServices() (output []byte, err error) {
	if output, err = runCompose("config", "--services"); err != nil {
		return
	}
	services = make(map[string]*Container)
	output = bytes.TrimSpace(output)
	for _, svr := range bytes.Split(output, []byte("\n")) {
		if output, err = runCompose("ps", "-q", string(svr)); err != nil {
			return
		}
		var (
			id   = string(bytes.TrimSpace(output))
			args = []string{"inspect", id, "--format", "'{{json .}}'"}
		)
		if output, err = exec.Command("docker", args...).CombinedOutput(); err != nil {
			logging.Warnf("exec.Command(docker) args(%v) stdout(%s) error(%v)", args, string(output), err)
			return
		}
		if output = bytes.TrimSpace(output); bytes.Equal(output, []byte("")) {
			err = fmt.Errorf("service: %s | container: %s fails to launch", svr, id)
			logging.Warnf("exec.Command(docker) args(%v) error(%v)", args, err)
			return
		}
		var c = &Container{}
		if err = json.Unmarshal(bytes.Trim(output, "'"), c); err != nil {
			logging.Warnf("json.Unmarshal(%s) error(%v)", string(output), err)
			return
		}
		services[string(svr)] = c
	}
	return
}
func getPort(k string) string {
	c, _ := services[k]
	for proto, ports := range c.NetworkSettings.Ports {
		if !strings.Contains(proto, "tcp") {
			continue
		}
		for _, publish := range ports {
			if publish.HostPort != "" {
				return publish.HostPort
			}
		}
	}
	return ""
}
func checkServices() (output []byte, err error) {
	var retry int
	defer func() {
		if err != nil && retry < 4 {
			retry++
			getServices()
			time.Sleep(time.Second * 2)
			output, err = checkServices()
			return
		}
		retry = 0
	}()
	for svr, c := range services {
		if err = c.Healthcheck(); err != nil {
			logging.Warnf("healthcheck(%s) error(%v) retrying %d times...", svr, err, 5-retry)
			return
		}
	}
	return
}

var healthchecks = map[string]func(*Container) error{"mysql": checkMysql, "mariadb": checkMysql}

// Healthcheck check container health.
func (c *Container) Healthcheck() (err error) {
	if status, health := c.State.Status, c.State.Health.Status; !c.State.Running || (health != "" && health != "healthy") {
		err = fmt.Errorf("service: %s | container: %s not running", c.GetImage(), c.GetID())
		logging.Warnf("docker status(%s) health(%s) error(%v)", status, health, err)
		return
	}
	if check, ok := healthchecks[c.GetImage()]; ok {
		err = check(c)
		return
	}
	for proto, ports := range c.NetworkSettings.Ports {
		if id := c.GetID(); !strings.Contains(proto, "tcp") {
			logging.Warnf("container: %s proto(%s) unsupported.", id, proto)
			continue
		}
		for _, publish := range ports {
			var (
				ip      = net.ParseIP(publish.HostIP)
				port, _ = strconv.Atoi(publish.HostPort)
				tcpAddr = &net.TCPAddr{IP: ip, Port: port}
				tcpConn *net.TCPConn
			)
			if tcpConn, err = net.DialTCP("tcp", nil, tcpAddr); err != nil {
				logging.Warnf("net.DialTCP(%s:%s) error(%v)", publish.HostIP, publish.HostPort, err)
				return
			}
			logging.Warnf("service(%s) health check,net.DialTCP(%s:%s) success", c.GetImage(), publish.HostIP, publish.HostPort)
			tcpConn.Close()
		}
	}
	return
}

func checkMysql(c *Container) (err error) {
	var ip, port, user, passwd string
	for _, env := range c.Config.Env {
		splits := strings.Split(env, "=")
		if strings.Contains(splits[0], "MYSQL_ROOT_PASSWORD") {
			user, passwd = "root", splits[1]
			continue
		}
		if strings.Contains(splits[0], "MYSQL_ALLOW_EMPTY_PASSWORD") {
			user, passwd = "root", ""
			continue
		}
		if strings.Contains(splits[0], "MYSQL_USER") {
			user = splits[1]
			continue
		}
		if strings.Contains(splits[0], "MYSQL_PASSWORD") {
			passwd = splits[1]
			continue
		}
	}
	var db *sql.DB
	if ports, ok := c.NetworkSettings.Ports["3306/tcp"]; ok {
		ip, port = ports[0].HostIP, ports[0].HostPort
	}
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, passwd, ip, port)
	if db, err = sql.Open("mysql", dsn); err != nil {
		logging.Warnf("sql.Open(mysql) dsn(%s) error(%v)", dsn, err)
		return
	}
	if err = db.Ping(); err != nil {
		logging.Warnf("ping(db) dsn(%s) error(%v)", dsn, err)
	}
	logging.Warnf("service(%s) health check,net.DialTCP(%s:%s) success", c.GetImage(), ip, port)
	defer db.Close()
	return
}
