package github_hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var log *logrus.Logger

var logLevels []*string

const logName = "logs/app.log"
const configPath = "scripts/config.json"
const scriptRoot = "scripts"

var appPath = "."

/**
这里负责程序的大部分初始化
1. 初始化日志目录以及文件
2. 初始化脚本目录以及文件配置
*/
func init() {
	initLog()
	err := initScriptConfig()
	if err != nil {
		panic(err)
	}
}

func initLog() {
	// 初始化日志配置
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.AddHook(CustomHook(fmt.Sprintf("%s/%s", appPath, logName), "0664"))
}

func initScriptConfig() error {
	// 初始化脚本配置
	err := ReadConfig(fmt.Sprintf("%s/%s", appPath, configPath))
	return err
}

// ReadConfig 读取脚本配置
func ReadConfig(fp string) error {
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	err = scriptConf.Flash(b)
	return err
}
