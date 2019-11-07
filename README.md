## github hook server
一个用来接收webhook的server

### 环境

golang v1.8+

### 安装

执行:
```bash
go get -u github.com/gaopengfei123123/hook
```

创建文件 `main.go`
```go
package main

import (
    "github.com/gaopengfei123123/hook"
)

func main() {
    hook.Execute()
}

```

执行命令
```bash
go build -o app main.go
```

`app` 这个二进制文件就是本体了, 放到合适的地方, 执行:
```bash
./app start -d
```
后台启动, app 同目录下会创建 `scripts`, `logs` 两个目录, 和一个 `hook.pid` 保存pid

检测是否启动:
```cmd
[root@xxx]# curl localhost:8080/ping
{"message":"pong v4"}
```
说明服务已经启动成功, 服务地址为`0.0.0.0:8080`




### 目录功能:
* logs 存放请求日志
* scripts 存放hook脚本, 当有对应库名的请求进来, 将执行配置好的脚本
* hook.pid 存放进程pid


TODO:
1. 平滑重启(done)
2. 调整配置文件结构, 支持scret验证

