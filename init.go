package hook

import (
	// "flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var log *logrus.Logger

var logLevels []*string

const logName = "logs/app.log"
const configPath = "scripts/config.json"
const scriptRoot = "scripts"

var appPath = "."
var appName = "hook"

/**
这里负责程序的大部分初始化
1. 初始化日志目录以及文件
2. 初始化脚本目录以及文件配置
*/
func init() {
	// 获取当前目录
	InitContextPath()
	// 初始化日志目录以及文件
	initLog()
	// 初始化脚本配置文件
	err := initScriptConfig()
	if err != nil {
		panic(err)
	}
}

func initLog() {
	// 初始化日志配置
	log = logrus.New()
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.AddHook(CustomHook(fmt.Sprintf("%s/%s", appPath, logName), "0664"))
}

func CmdAndProgress() {
	command := exec.Command("pwd")
	command.Start()
	fmt.Printf("hook start, [PID] %d running...\n", command.Process.Pid)
	ioutil.WriteFile("hook.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
}

// 初始化基础配置(这个函数一定要最先执行)
func InitContextPath() {
	pt := GetCurrPath()
	currnetPath := filepath.Dir(pt)
	currentDir := filepath.Base(pt)
	// fmt.Printf("当前文件夹: %s/%s \n", currnetPath, currentDir)
	appPath = fmt.Sprintf("%s/%s", currnetPath, currentDir)
	appName = filepath.Base(os.Args[0])
}

// 获取当前目录
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

// 初始化脚本目录
func initScriptConfig() error {
	confFullPath := fmt.Sprintf("%s/%s", appPath, configPath)

	// 检测目录是否存在
	dirpath, _ := path.Split(confFullPath)
	if !IsExist(dirpath) {
		dirPerm, _ := strconv.ParseInt("0755", 8, 64)
		err := os.MkdirAll(dirpath, os.FileMode(dirPerm))
		if err != nil {
			return err
		}
		// log.Infof("初始化脚本目录, 路径为: %s", dirpath)
	}

	if !IsExist(confFullPath) {
		filePerm, _ := strconv.ParseInt("0744", 8, 64)
		fd, err := os.OpenFile(confFullPath, os.O_WRONLY|os.O_CREATE, os.FileMode(filePerm))
		if err != nil {
			return err
		}
		_, err = fd.Write([]byte(`{"repository-name":"script-name"}`))
		if err != nil {
			return err
		}
		fd.Close()
		// log.Infof("初始化脚本文件, 路径为: %s", confFullPath)
	}

	// 初始化脚本配置
	err := ReadConfig(confFullPath)
	return err
}

// IsExist 文件或目录是否存在
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsExist(err) {
		return true
	}
	return false
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
