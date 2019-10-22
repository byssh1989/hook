package github_hook

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/push", PushHookHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func PushHookHandler(c *gin.Context) {
	data, _ := c.GetRawData()
	// fmt.Printf("%s \n", data)

	params := GITHUB_HOOK{}
	json.Unmarshal(data, &params)

	fmt.Printf("输出: %v \n", params)
}

type GITHUB_HOOK struct {
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
