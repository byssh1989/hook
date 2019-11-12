package hook

import (
	"encoding/json"
	"fmt"
	"sync"
)

var scriptConf = scriptConfig{}

// ScriptConfig 库与脚本文件名的映射关系
type scriptConfig struct {
	sync.RWMutex
	data    bashMap
	repoMap map[string]repo
}
type bashMap map[string]string
type repo struct {
	Secret     string            `json:"secret"`
	ScriptPath string            `json:"script_path"`
	Event      map[string]string `json:"event"`
}

func (c *scriptConfig) Get(key string) (conf repo, err error) {
	c.RLock()
	defer c.RUnlock()

	conf, ok := c.repoMap[key]
	if !ok {
		err = fmt.Errorf("找不到对应command, key: %s", key)
	}
	return
}

// Set 目前计划不允许重复设置
func (c *scriptConfig) Set(key string, val repo) (conf repo, err error) {
	c.Lock()
	defer c.Unlock()

	conf, ok := c.repoMap[key]
	if !ok {
		c.repoMap[key] = val
		conf = val
	} else {
		err = fmt.Errorf("当前库 %s 已经存在config", key)
	}
	return
}

// Flash 重新加载配置
func (c *scriptConfig) Flash(data []byte) (err error) {
	c.Lock()
	defer c.Unlock()

	tmp := map[string]repo{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	c.repoMap = tmp
	return
}

// 获取密钥
func (r repo) ValidateSign(payload []byte, sign string) (err error) {
	if r.Secret == "" {
		return
	}
	pass := checkSecret(payload, r.Secret, sign)
	if !pass {
		err = ErrSignature
	}
	return
}

func (r repo) EventBash(event string) (cmd string, err error) {
	pt := r.ScriptPath
	if pt == "" {
		pt = fmt.Sprintf("%s/%s", appPath, scriptRoot)
	}

	command, ok := r.Event[event]
	if !ok {
		err = fmt.Errorf("[Event: %v] 不存在", event)
	}

	cmd = fmt.Sprintf("%s/%s", pt, command)
	return
}
