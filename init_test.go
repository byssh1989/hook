package github_hook

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestHook(t *testing.T) {
	log.Info("测试logrus info")
	log.Warn("测试logrus warn")
	log.Debug("测试logrus debug")
	log.WithFields(logrus.Fields{
		"var": "params",
	}).Warn("测试logrus warn")

}
