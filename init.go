package github_hook

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var logLevels []*string

const logName = "Users/gpf/Documents/www/go_projects/src/app.log"

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.AddHook(CustomHook("./logs/app.log", "0664"))
}
