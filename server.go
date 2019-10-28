package github_hook

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os/exec"
)

// Start 启动服务
func Start() {
	r := gin.Default()
	log.Info("start...")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/push", PushHookHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}

// PushHookHandler 处理推送事件
func PushHookHandler(c *gin.Context) {
	data, _ := c.GetRawData()
	params := GithubHook{}
	json.Unmarshal(data, &params)

	fields := logrus.Fields{}
	json.Unmarshal(data, &fields)
	log.WithFields(fields).Info("request_raw")

	// fmt.Printf("%s \n", data)

	err := execHookBash(params)
	if err != nil {
		log.Error(err)

		c.JSON(401, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	}

}

// GithubHook github的json结构
type GithubHook struct {
	Zen    string
	HookID int
	Hook   struct {
		Type   string
		ID     int
		Name   string
		Active bool
		Events []string
	}
	Repository struct {
		ID       int
		NodeID   string `json:"node_id"`
		Name     string
		FullName string `json:"full_name"`
		Private  bool
	}
	Sender struct {
		ID        int
		Login     string
		NodeID    string
		Type      string
		SiteAdmin bool
	}
}

// 项目名映射得脚本文件
var configs = map[string]string{
	"github_hook": "./scripts/hook",
}

// 执行脚本目录
const hookScriptDirPath = "./scripts"

// 执行推送的逻辑
func execHookBash(hook GithubHook) error {

	command, ok := configs[hook.Repository.Name]
	command = fmt.Sprintf("%s/%s", hookScriptDirPath, command)
	log.Infof("Execute command: %s", command)

	if !ok {
		return fmt.Errorf("config中找不到相应库得配置, configs: %v, name: %s, hook: %v", configs, command, hook)
	}

	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()

	if err != nil {
		log.Errorf("Execute Shell:%s failed with error:%s", command, err.Error())
		return fmt.Errorf("Execute Shell:%s failed with error:%s", command, err.Error())
	}
	log.Infof("Execute Shell:%s finished with output:\n%s", command, string(output))
	return nil
}
