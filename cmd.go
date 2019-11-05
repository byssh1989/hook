package github_hook

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

var daemon bool
var rootCmd = &cobra.Command{
	Use:   "hook",
	Short: "一个钩子服务器",
	Long:  `接受github钩子请求, 执行对应脚本`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("啦啦啦~")
		fmt.Println("请执行 hook -h 查看具体命令")
	},
}

func Execute() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "启动命令",
		Run: func(cmd *cobra.Command, args []string) {
			if !daemon {
				fmt.Println("前台启动")
				pid := os.Getpid()
				fmt.Printf("getpid: %d \n", pid)
				setPidFile(pid)
				Start()
			} else {
				fmt.Println("后台启动")
				appCmd := fmt.Sprintf("%s/%s", appPath, appName)
				command := exec.Command(appCmd, "start")
				err := command.Start()
				if err != nil {
					panic(err)
				}
				fmt.Printf("hook start, [PID] %d running...\n", command.Process.Pid)
				setPidFile(command.Process.Pid)
			}
		},
	}
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "终止命令",
		Run: func(cmd *cobra.Command, args []string) {
			Stop()
		},
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "程序版本",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("当前版本 v0.1.1")
		},
	}

	reloadCmd := &cobra.Command{
		Use:   "reload",
		Short: "重新加载",
		Run: func(cmd *cobra.Command, args []string) {
			Reload()
		},
	}

	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	startCmd.Flags().BoolVarP(&graceful, "graceful", "g", false, "is graceful restart?")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(reloadCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 获取程序pid值
func getPid() int {
	pidPath := fmt.Sprintf("%s/%s", appPath, "hook.pid")
	b, err := ioutil.ReadFile(pidPath)

	if err != nil {
		if err == os.ErrNotExist {
			fmt.Println("pid 不存在")
			os.Exit(0)
		} else {
			panic(err)
		}
	}
	pid, _ := strconv.Atoi(fmt.Sprintf("%s", b))
	return pid
}

// 移除pid文件
func removePidFile() {
	pidPath := fmt.Sprintf("%s/%s", appPath, "hook.pid")
	err := os.Remove(pidPath)
	if err != nil {
		panic(err)
	}
}

func setPidFile(pid int) {
	ioutil.WriteFile(fmt.Sprintf("%s/%s", appPath, "hook.pid"), []byte(fmt.Sprintf("%d", pid)), 0666)
}

// Stop 终止后台进程
func Stop() {
	pid := getPid()
	pro, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	err = pro.Signal(os.Interrupt)
	if err != nil {
		panic(err)
	}
	removePidFile()

	fmt.Println("程序已退出, bye")
}

// 向后台发送信号, 重载进程
func Reload() {
	pid := getPid()

	pro, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	err = pro.Signal(syscall.SIGUSR2)
	if err != nil {
		panic(err)
	}

	fmt.Println("程序已重启")
}
