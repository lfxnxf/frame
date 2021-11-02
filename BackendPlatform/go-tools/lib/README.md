# 通用用户资料获取
批量分片获取用户资料

```go
uids := []int64{218313629, 564238347}
start := 200000
count := 1000
for i := start; i < start+count; i++ {
    uids = append(uids, int64(i))
}
now:=time.Now()
userinfos, err := GetUserInfo(context.TODO(), uids)
```