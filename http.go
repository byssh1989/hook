package hook

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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

	r.POST("/push", GithubSecret(), PushHookHandler)
	GraceRun(":8080", r) // listen and serve on 0.0.0.0:8080
	// r.Run()
}

// PushHookHandler 处理推送事件
func PushHookHandler(c *gin.Context) {

	params, err := InitGithubHook(c)

	if err != nil {
		log.Error(err)
		switch err {
		case ErrTimeout:
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		case ErrSignature:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	cmd, err := selectCMDByHook(params)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = SendTask(cmd)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
}

func InitGithubHook(c *gin.Context) (hk GithubHook, err error) {
	data, _ := c.GetRawData()
	sign := c.GetHeader("X-Hub-Signature")
	event := c.GetHeader("X-GitHub-Event")

	fields := logrus.Fields{}
	fields["raw"] = string(data)
	log.WithFields(fields).Info("request_raw")

	hk = GithubHook{}
	err = json.Unmarshal(data, &hk)
	if err != nil {
		return
	}

	hk.Event = event
	hk.Signature = sign
	hk.Payload = data
	return
}
