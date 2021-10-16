package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of Hugo",
	Long: "All software has versions. This id Hugo's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 --HEAD")
	},
}

func init() {
	// 通过调用 rootCmd.AddCommand(versionCmd) 给 rootCmd 命令
	// 添加了一个 versionCmd 命令。
	rootCmd.AddCommand(versionCmd)
}
