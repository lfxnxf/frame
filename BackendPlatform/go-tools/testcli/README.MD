## testcli UT运行环境构建工具
基于 docker-compose 实现跨平台跨语言环境的容器依赖管理方案。  
工具通过读取用户config.toml配置，自动读取用户依赖的mysql、redis、kafka等配置，启动对应的容器提供用户使用。

### 功能和特性
- 读取config.toml，自动启动资源容器提供使用
- 自动清理启动资源


### 编译安装
*使用本工具/库需要前置安装好 docker & docker-compose@v1.24.1^*



### 初始化数据
由于服务读取config.toml，根据配置文件生成对应资源，所以对于mysql需要提供建表等初始化sql（放在config.toml同级目录），对于kafka需要提供对应topic，工具会在容器启动时自动初始化各个容器
目录结构
```text
├── ali-test
│   ├── config.toml
│   └── shopping_store.sql

```
kafka生产者配置
```text
[[kafka_producer_client]]
    producer_to = "busi.privilege.new_shopping-kafka"
    kafka_broken="ali-a-inf-kafka-test11.bj:9092,ali-c-inf-kafka-test12.bj:9092,ali-a-inf-kafka-test13.bj:9092"
    topic="topic1,topic2" # 填写初始化的kafka
    retrymax=1
    get_error = true
    get_success = true
    required_acks="WaitForAll"
    use_sync=true
```



### 运行


example:
```Go
package dao

import (
	"os"
	"testing"

	"github.com/lfxnxf/frame/BackendPlatform/go-tools/testcli"
	"github.com/lfxnxf/frame/bpc/server/busi/busi.privilege.new_shopping/conf"
	"github.com/lfxnxf/frame/logic/daenerys"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	path := testcli.Setup("../app/config/ali-test/config.toml") // 启动配置对应的容器资源
	// defer testcli.Teardown() // 移除容器资源
	daenerys.Init(daenerys.ConfigPath(path))                    // 框架初始化
	conf, err := conf.Init()
	if err != nil {
		return
	}
	d = New(conf)
	os.Exit(m.Run())
}

 ```
 运行测试case
 ```text
go test -p 1 -v  ./...
```
如果存在多个目录都需要依赖容器资源，可以不用关闭容器资源，自己手动关闭资源
```text
docker stop $(docker ps -aq) && docker rm $(docker ps -aq)
```
## 注意
因为启动mysql容器较为缓慢，需要稍等一会儿