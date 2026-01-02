package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-pal",
	Short: "Your friendly AI-powered development assistant.",
	Long: `dev-pal is a CLI tool that uses AI to help with common development tasks
like writing commit messages, generating PR descriptions, and more.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
