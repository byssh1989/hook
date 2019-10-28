package github_hook

import (
	// "github.com/lestrrat-go/file-rotatelogs"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strconv"
)

// CustomHook 自定义hook
func CustomHook(filename, perm string) logrus.Hook {
	h := Hook{}
	h.Filename = filename
	h.Perm = perm
	h.init()
	return h
}

// Hook 日志钩子结构体
type Hook struct {
	Filename string
	Perm     string
	Fd       *os.File
}

// Levels 日志可用级别
func (h Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 日志写入函数
func (h Hook) Fire(f *logrus.Entry) error {
	serialized, err := f.Logger.Formatter.Format(f)
	if err != nil {
		return err
	}
	_, err = h.Fd.Write([]byte(serialized))
	return err
}

// hook接口初始化, 创建相关文件
func (h *Hook) init() {
	err := h.createFile()
	if err != nil {
		panic(fmt.Errorf("日志初始化错误, Error: %v", err))
	}
}

func (h *Hook) createFile() error {
	perm, err := strconv.ParseInt(h.Perm, 8, 64)
	if err != nil {
		return err
	}

	dirpath := path.Dir(h.Filename)
	err = os.MkdirAll(dirpath, os.FileMode(perm))
	if err != nil {
		return err
	}

	filename := h.Filename
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(perm))
	h.Fd = fd
	return err
}
