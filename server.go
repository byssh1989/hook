package github_hook

import (
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
	fmt.Printf("%s \n", data)

}
