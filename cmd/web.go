package cmd

import (
	gowebtemplate "github.com/menggggggg/go-web-template/cmd/go-web-template"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run web service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gowebtemplate.RunGoWebTemplateServer()
	},
}
