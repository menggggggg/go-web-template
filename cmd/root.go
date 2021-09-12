package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "web",
		Short: "golang web project template",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// 默认启动web项目
			webCmd.Run(cmd, args)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
