package github_hook

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
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
				ioutil.WriteFile(fmt.Sprintf("%s/%s", appPath, "hook.pid"), []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
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

	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(stopCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Stop 终止后台进程
func Stop() {
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
	pro, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}
	err = pro.Signal(os.Interrupt)

	if err != nil {
		panic(err)
	}

	err = os.Remove(pidPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("hook 已退出, bye")
}
