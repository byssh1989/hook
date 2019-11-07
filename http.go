package hook

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

	r.POST("/push", LogMiddleware(), PushHookHandler)
	GraceRun(":8080", r) // listen and serve on 0.0.0.0:8080
	// r.Run()
}

// PushHookHandler 处理推送事件
func PushHookHandler(c *gin.Context) {
	data, _ := c.GetRawData()
	salt := "123123"
	sign := c.GetHeader("X-Hub-Signature")

	if !checkSecret(data, salt, sign) {
		c.JSON(401, gin.H{
			"error": "签名错误",
		})
		return
	}

	params := GithubHook{}
	json.Unmarshal(data, &params)

	cmd, err := selectCMDByHook(params)
	if err != nil {
		log.Error(err)
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = SendTask(cmd)
	if err != nil {
		log.Error(err)
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
		})
		return
	}
}
