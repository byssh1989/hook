## github hook server
一个用来接收webhook的server


### 目录功能:
* logs 存放请求日志
* scripts 存放hook脚本, 当有对应库名的请求进来, 将执行配置好的脚本
* hook.pid 存放进程pid


TODO:
1. 平滑重启(done)
2. 调整配置文件结构, 支持scret验证

