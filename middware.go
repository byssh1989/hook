package hook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		event := c.GetHeader("X-GitHub-Event")
		fmt.Printf("X-GitHub-Event: %v \n", event)

		data, _ := c.GetRawData()
		fields := logrus.Fields{}
		fields["raw"] = string(data)
		log.WithFields(fields).Info("request_raw")

		c.Next()
	}
}
