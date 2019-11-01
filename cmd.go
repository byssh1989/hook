package github_hook

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var daemon bool
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务",
	Long:  `启动一个web服务器, 接收hook请求`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("啦啦啦")
	},
}

func Execute() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "启动命令",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				fmt.Println("后台启动")
			} else {
				fmt.Println("前台启动")
				Start()
			}
			fmt.Printf("启动 %v \n", daemon)
		},
	}
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "终止命令",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("终止 %v \n", daemon)
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
