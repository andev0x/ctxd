package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lea",
	Short: "lea is a structural context engine for AI-native engineering",
	Long:  `lea helps AI models and developers understand large codebases by modeling repositories as living structural graphs.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Root flags if any
}
