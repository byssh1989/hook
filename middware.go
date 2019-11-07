package hook

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GithubSecret() gin.HandlerFunc {
	return func(c *gin.Context) {
		event := c.GetHeader("X-GitHub-Event")
		fmt.Printf("X-GitHub-Event: %v \n", event)
		c.Next()
	}
}
