package hook

import (
	"bufio"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
	"os/exec"
	"time"
)

var cmdChan chan string

func init() {
	cmdChan = make(chan string, 1000)
	go StartCmdQuene()
}

// GithubHook github的json结构
type GithubHook struct {
	Event     string `json:"-"`
	Signature string `json:"-"`
	Payload   []byte `json:"-"`
	Zen       string
	HookID    int
	Hook      struct {
		Type   string
		ID     int
		Name   string
		Active bool
		Events []string
	}
	Repository struct {
		ID       int
		NodeID   string `json:"node_id"`
		Name     string
		FullName string `json:"full_name"`
		Private  bool
		PushedAt int64 `json:"pushed_at"`
	}
	Sender struct {
		ID        int
		Login     string
		NodeID    string
		Type      string
		SiteAdmin bool
	}
	HeadCommit headCommit `json:"head_commit"`
}

type headCommit struct {
	ID                       string
	Timestamp                string
	Added, Removed, Modified []string
}

// StartCmdQuene 启动命令行队列
func StartCmdQuene() {
	log.Info("启动queue")
	for {
		task := <-cmdChan
		log.Infof("收到命令: %s \n", task)
		go func(task string) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
			ch := make(chan int)
			defer close(ch)

			go func(ch chan int, cmd string) {
				err := execBash(cmd)
				if err != nil {
					log.Errorf("执行任务出错, cmd: %s; err: %v \n", task, err)
				}
				ch <- 1
			}(ch, task)

		LOOP:
			for {
				select {
				case <-ctx.Done():
					log.Errorf("执行任务超时, cmd: %s \n", task)
				case <-ch:
					cancel()
					break LOOP
				}
			}
		}(task)
	}
}

// SendTask 发送要执行的命令
func SendTask(task string) error {
	if cap(cmdChan) > 1000 {
		return fmt.Errorf("消息堆积, 请稍候再试")
	}
	cmdChan <- task
	return nil
}

// 提取对应的cmd
func selectCMDByHook(hook GithubHook) (command string, err error) {
	// 规避一下重放问题
	now := time.Now().Unix()
	cha := now - hook.Repository.PushedAt
	if cha > 30 || cha < -30 {
		err = ErrTimeout
		return
	}

	// 把脚本的路径拼一下
	conf, err := scriptConf.Get(hook.Repository.Name)
	if err != nil {
		return
	}

	if err = conf.ValidateSign(hook.Payload, hook.Signature); err != nil {
		return "", err
	}

	command, err = conf.EventBash(hook.Event)
	log.Infof("Execute command: %s", command)
	if err != nil {
		return
	}

	return
}

// 执行推送的逻辑
func execBash(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	// 非阻塞输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorf("Execute Shell:%s failed with error:%s", command, err.Error())
		return fmt.Errorf("Execute Shell:%s failed with error:%s", command, err.Error())
	}

	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		log.Infof("%s", line)
	}
	cmd.Wait()

	log.Infof("Execute Shell:%s finished", command)
	return nil
}

func checkSecret(payload []byte, salt, sign string) bool {
	// log.Infof("payload: %x, salt: %s, sign: %s \n", payload, salt, sign)
	mac := hmac.New(sha1.New, []byte(salt))
	mac.Write(payload)
	res := fmt.Sprintf("sha1=%x", mac.Sum(nil))
	return res == sign
}
