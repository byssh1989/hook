## clone from https://github.com/gaopengfei123123/hook

克隆自https://github.com/gaopengfei123123/hook

## github hook server

一个用来接收 webhook 的 server

### 环境

golang v1.11+ (因为用到了 go mod)

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

后台启动, app 同目录下会创建 `scripts`, `logs` 两个目录, 和一个 `hook.pid` 保存 pid

检测是否启动:

```cmd
[root@xxx]# curl localhost:8080/ping
{"message":"pong v5"}
```

说明服务已经启动成功, 服务地址为`0.0.0.0:8080`

### 接口

- `0.0.0.0:8080/ping` 检测接口
- `0.0.0.0:8080/push` 接受 github webhook 请求, 根据 Repository.Name 去判断执行什么脚本

### 可用指令

```
reload      重新加载日志, 以及平滑重启
start       启动命令, -d 后台运行
stop        终止命令
version     程序版本
```

### 目录功能:

- logs 存放请求日志
- scripts 存放 hook 脚本, 当有对应库名的请求进来, 将执行配置好的脚本
- hook.pid 存放进程 pid

### 配置文件

```json
{
  "github_hook": {
    // 以请求消息中的 repository.name 字段来做key
    "secret": "xxxxxx", // 如果设置了secret则会用这个进行验证, 为空则不验证
    "script_path": "", // 脚本所在绝对目录, 为空就是当前的script目录
    "event": {
      "push": "test" // 推送事件执行的脚本
    }
  }
}
```

### 特性

- 一键初始化
- 平滑重启, 信号通信
- 支持 secret 验证
- 异步执行脚本
- 规避重放
- 脚本传递参数(TODO)
- 事件钩子(TODO)

### 依赖组件

- http 框架 [gin](https://github.com/gin-gonic/gin)
- 日志组件 [logrus](https://github.com/sirupsen/logrus)
- 命令行组件 [cobra](https://github.com/spf13/cobra)
