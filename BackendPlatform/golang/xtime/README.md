# 时间帮助相关
### xtime.Time使用 
数据库时间建议都使用TIMESTAMP，传输给客户端建议统一用时间戳（优点：统一、规范、无时区概念）
 
```text
type MessageTable struct {
	ID         int64      `gorm:"column:id" json:"id"`
	Msg        string     `gorm:"column:desc" json:"desc"`
	UpdateTime xtime.Time `gorm:"column:update_time" json:"update_time"` // 数据库类型 TIMESTAMP
}
// marshal之后时间会转为int
如：
{"id":1,"msg":"a msg","update_time":1577934399}

```

### xtime.Duration使用
配置项duration用字符串表示更加直观    
支持单位：小时h、分钟m、秒s、毫秒ms、微秒us、纳秒ns  
未列出的就是不支持单位  
```text
// 配置如下
[demo]
    idleCount=100
    execTimeout = "300ms"
    tranTimeout = "400ms"


type Demo struct {
    IdleCount int   `toml:"idleCount"`
    ExecTimeout  xtime.Duration `toml:"execTimeout"`
    TranTimeout  xtime.Duration `toml:"tranTimeout"`
}
// 对应结构体
type Conf struct {
    Demo *Demo `toml:"demo"`
}

// 使用时转为标准time.Duration使用:
time.Duration(x.ExecTimeout)

```