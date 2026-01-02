package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var baseBranch string

var generatePrDescCmd = &cobra.Command{
	Use:   "generate-pr-desc",
	Short: "Generates a PR description for the current branch.",
	Long:  `Analyzes the commits and diff of the current branch against a base branch (default 'main') and uses AI to generate a PR title and description.`,
	Run: func(cmd *cobra.Command, args []string) {
		description, err := generatePRDescription(baseBranch)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ AI has generated your PR description:")
		fmt.Println("------------------------------------------------")
		fmt.Println(description)
		fmt.Println("------------------------------------------------")
		fmt.Println("Copy the text above and use it for your Pull Request.")
	},
}

func init() {
	generatePrDescCmd.Flags().StringVar(&baseBranch, "base", "main", "The base branch to compare against.")
	rootCmd.AddCommand(generatePrDescCmd)
}
