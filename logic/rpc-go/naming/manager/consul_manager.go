package manager

import (
	"errors"
	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"sync"
	"time"
	// consul "github.com/lfxnxf/frame/inkelogic/rpc-go/discovery/consul"
	"github.com/lfxnxf/frame/BackendPlatform/golang/discovery/consul"
	naming "github.com/lfxnxf/frame/logic/rpc-go/naming"
)

var (
	CONSUL_AGENT_ADDR = "127.0.0.1:8500"
)

var (
	CLIENT_NOT_INIT_OK = errors.New("client not init ok or bad")
)

type WatchInfo struct {
	lastIndex uint64
	lastAddrs []string
}

type KVInfo struct {
	lastIndex uint64
	lastvalue string
}

type ConsuleManager struct {
	client    *consul.Client
	close     chan bool
	watchinfo map[string]WatchInfo
	kvinfo    map[string]KVInfo
	messages  chan *naming.Message
}

var clientLock *sync.Mutex
var lock *sync.Mutex
var kvlock *sync.Mutex

func init() {
	lock = new(sync.Mutex)
	kvlock = new(sync.Mutex)
	clientLock = new(sync.Mutex)
}

func NewConsuleManager() *ConsuleManager {

	clientmanager := &ConsuleManager{
		watchinfo: make(map[string]WatchInfo),
		kvinfo:    make(map[string]KVInfo),
		close:     make(chan bool),
		messages:  make(chan *naming.Message),
	}
	return clientmanager
}

func (cm *ConsuleManager) createClient() error {

	clientLock.Lock()
	defer clientLock.Unlock()

	var consulAgent []string
	consulAgent = append(consulAgent, CONSUL_AGENT_ADDR)
	c, err := consul.New(consulAgent, "http", nil)
	if err != nil {
		log.BalanceLog("new consul agent error,addr:", CONSUL_AGENT_ADDR)
		return err
	}
	cm.client = c
	return nil
}

func (cm *ConsuleManager) Start() error {
	lock.Lock()
	defer lock.Unlock()
	return cm.createClient()
}

func (cm *ConsuleManager) checkClient() error {

	clientLock.Lock()
	defer clientLock.Unlock()

	if cm.client == nil {
		var consulAgent []string
		consulAgent = append(consulAgent, CONSUL_AGENT_ADDR)
		c, err := consul.New(consulAgent, "http", nil)
		if err != nil {
			return CLIENT_NOT_INIT_OK
		}
		cm.client = c
		return nil
	}
	return nil
}

func (cm *ConsuleManager) UnRegister() error {
	errc := cm.checkClient()
	if errc != nil {
		return errc
	}

	if cm.close != nil {
		close(cm.close)
		cm.close = nil
	}
	return nil
}

func (cm *ConsuleManager) RegisterService(targets []string, proto string, tags []string, address string, port int) error {

	errc := cm.checkClient()
	if errc != nil {
		return errc
	}

	err := cm.client.RegisterMultTagService(targets, proto, tags, address, port, cm.close)
	if err != nil {
		return err
	}
	return nil
}

func makeService(target, proto string) string {
	if proto == "http" {
		target += "-http"
	} else {
		target += "-" + proto
	}
	return target
}

func (cm *ConsuleManager) GetService(target, proto, tag, dc string) ([]string, error) {

	errc := cm.checkClient()
	if errc != nil {
		var a []string
		return a, errc
	}

	lock.Lock()
	defer lock.Unlock()

	addr, index, err := cm.client.GetServiceByDc(target, proto, tag, dc)

	if err != nil {
		return addr, err
	}

	if info, ok := cm.watchinfo[target]; ok {
		info.lastIndex = index
		cm.watchinfo[makeService(target, proto)] = info
	} else {
		info := WatchInfo{
			lastIndex: index,
		}
		cm.watchinfo[makeService(target, proto)] = info
	}
	log.BalanceLog("consul getservice,service:", makeService(target, proto))
	return addr, err
}

func (cm *ConsuleManager) GetKvValue(path string) (string, error) {
	key := path
	value, _, err := cm.client.GetKeyValue(key)
	return value, err
}

func (cm *ConsuleManager) WatchKV(path string) {
	go cm.watchkv(path)
}

func (cm *ConsuleManager) Watch(target, proto, tag, dc string) {
	go cm.watch(target, proto, tag, dc)
}

func (cm *ConsuleManager) GetValues(keys []string) (map[string]string, error) {
	var info map[string]string
	info = make(map[string]string)
	return info, nil
}

func checkInArr(key string, arr []string) bool {
	for _, a := range arr {
		if key == a {
			return true
		}
	}
	return false
}

func addrIsChange(arr1 []string, arr2 []string) bool {

	for _, a1 := range arr1 {
		flag := checkInArr(a1, arr2)
		if flag == false {
			return true
		}
	}

	for _, a2 := range arr2 {
		flag := checkInArr(a2, arr1)
		if flag == false {
			return true
		}
	}
	return false
}

func (cm ConsuleManager) checkAddrisChange(target, proto, tag string, address []string, lastIndex uint64) bool {

	lock.Lock()
	defer lock.Unlock()

	targetN := makeService(target, proto)

	if _, ok := cm.watchinfo[targetN]; !ok {
		info := WatchInfo{
			lastIndex: lastIndex,
		}
		cm.watchinfo[targetN] = info
	}
	info := cm.watchinfo[targetN]

	lastAddr := info.lastAddrs
	change := false

	change = addrIsChange(address, lastAddr)

	if change {
		info.lastAddrs = address
		log.BalanceLog("ip change,target:", target, ",proto:", proto, ",addr:", address, ",lastAddr:", lastAddr)
	} else {
		log.BalanceLog("ip no change,target:", target, ",proto:", proto, ",addr:", address, ",lastAddr:", lastAddr)
	}
	info.lastIndex = lastIndex
	cm.watchinfo[targetN] = info
	return change
}

func (cm *ConsuleManager) dowatchmsg(target, proto, tag string, address []string, lastIndex uint64) {

	change := cm.checkAddrisChange(target, proto, tag, address, lastIndex)
	if change {
		a := &naming.Message{Addrs: address, Target: target, Proto: proto, Type: naming.SERVICE}
		cm.messages <- a
	}
}

func (cm *ConsuleManager) watch(target, proto, tag, dc string) ([]string, error) {

	for {

		errc := cm.checkClient()
		if errc != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		address, lastIndex, err := cm.client.GetServiceByDc(target, proto, tag, dc)
		log.BalanceLogf("consul watcher get service %+v, err %v, target:%v,proto:%v,err:%v,lastIndex:%v\n", address, err, target, proto, err, lastIndex)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}
		cm.dowatchmsg(target, proto, tag, address, lastIndex)
		lastIndex, err = cm.client.WatchService(target, proto, tag, lastIndex, cm.close)
		if err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}

func (cm *ConsuleManager) checkValueIsChange(value string, lastvalue string) bool {

	if value == lastvalue {
		return false
	}
	return true
}

func (cm *ConsuleManager) checkKvChange(path string, value string, lastIndex uint64) bool {

	kvlock.Lock()
	defer kvlock.Unlock()

	key := path
	if _, ok := cm.kvinfo[key]; !ok {
		info := KVInfo{
			lastIndex: lastIndex,
		}
		cm.kvinfo[key] = info
	}
	v := cm.kvinfo[key]

	lastvalue := v.lastvalue

	change := false
	change = cm.checkValueIsChange(value, lastvalue)

	if change {
		v.lastvalue = value
		log.BalanceLog("kv changen,path:", key, ",last_value:", lastvalue, ",new_value:", value)
	} else {
		if len(value) != 0 {
			log.BalanceLog("kv no change,path:", key, ",last_value:", lastvalue, ",new_value:", value)
		}
	}

	v.lastIndex = lastIndex
	cm.kvinfo[key] = v

	return change
}

func (cm *ConsuleManager) dowatchkv(path string, value string, lastIndex uint64) {

	change := cm.checkKvChange(path, value, lastIndex)
	if change {
		a := &naming.Message{KValue: value, Type: naming.KV, Key: path}
		cm.messages <- a
	}
}

func (cm *ConsuleManager) watchkv(path string) {

	for {
		err := cm.checkClient()
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		key := path
		value, lastIndex, err := cm.client.GetKeyValue(key)

		if err != nil || len(value) == 0 {
			time.Sleep(15 * time.Second)
			continue
		}

		cm.dowatchkv(key, value, lastIndex)

		lastIndex, err = cm.client.WatchPrefix(key, lastIndex, cm.close)
		if err != nil {
			log.BalanceLog("watch kv msg error path:", path, ",lastIndex:", lastIndex, ",err:", err)
			time.Sleep(10 * time.Second)
		}

	}
}

func (cm *ConsuleManager) Next() (*naming.Message, error) {
	u := <-cm.messages
	return u, nil
}
