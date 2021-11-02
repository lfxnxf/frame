# 参与贡献

1.pull代码

2.修改代码

3.运行测试 ./check.sh

4.提交合并

5.review代码

6.OK 合并到master

```text
go test -race -cover $(go list ./... | grep -v "vendor")
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/cast       (cached)        coverage: 99.7% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/crypto     1.038s  coverage: 85.7% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/errs       1.065s  coverage: 86.4% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/ext        (cached)        coverage: 100.0% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/gid        1.048s  coverage: 80.9% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/http       (cached)        coverage: 0.0% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/lib        1.434s  coverage: 92.5% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/req        (cached)        coverage: 38.5% of statements
ok      github.com/lfxnxf/frame/BackendPlatform/go-tools/valid      (cached)        coverage: 66.7% of statements

```****