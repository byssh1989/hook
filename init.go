package github_hook

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var log *logrus.Logger

var logLevels []*string

const logName = "./logs/app.log"
const configPath = "./scripts/config.json"

func init() {
	// 初始化日志配置
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.AddHook(CustomHook(logName, "0664"))

	// 初始化脚本配置
	err := ReadConfig(configPath)
	if err != nil {
		panic(err)
	}
}

// ReadConfig 读取脚本配置
func ReadConfig(fp string) error {
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &configs)
	return err
}
