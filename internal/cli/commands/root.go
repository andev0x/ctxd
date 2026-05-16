package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ctxd",
	Short: "ctxd is a structural context engine for AI-native engineering",
	Long:  `ctxd helps AI models and developers understand large codebases by modeling repositories as living structural graphs.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Root flags if any
}
