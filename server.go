package hook

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
			"message": "pong v4",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello hook server",
		})
	})

	r.POST("/push", PushHookHandler)
	GraceRun(":8080", r) // listen and serve on 0.0.0.0:8080
	// r.Run()
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

// 执行推送的逻辑
func execHookBash(hook GithubHook) error {
	// 每次都读一下脚本配置
	err := initScriptConfig()
	if err != nil {
		return err
	}

	// 把脚本的路径拼一下
	command, err := scriptConf.Get(hook.Repository.Name)
	command = fmt.Sprintf("%s/%s/%s", appPath, scriptRoot, command)

	log.Infof("Execute command: %s", command)
	if err != nil {
		log.Errorf("找不到对应的command: %s, Err: %s", command, err.Error())
		return err
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
