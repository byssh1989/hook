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


### 接口

* `0.0.0.0:8080/ping`   检测接口
* `0.0.0.0:8080/push`   接受github webhook请求, 根据 Repository.Name 去判断执行什么脚本



### 目录功能:
* logs 存放请求日志
* scripts 存放hook脚本, 当有对应库名的请求进来, 将执行配置好的脚本
* hook.pid 存放进程pid


### 特性
* 一键初始化
* 平滑重启, 信号通信
* 支持secret验证
* 异步执行脚本
* 规避重放
* 脚本传递参数(TODO)
* 事件钩子(TODO)


