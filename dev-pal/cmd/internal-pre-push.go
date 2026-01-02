package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var remoteName, remoteUrl string

var internalPrePushCmd = &cobra.Command{
	Use:    "internal-pre-push",
	Short:  "Internal command for the pre-push hook.",
	Long:   `This command is not meant to be called directly by the user. It is used by the pre-push git hook to provide agentic functionality.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// This hook receives refs from stdin, one per line, in the format:
		// <local-ref> <local-sha> <remote-ref> <remote-sha>
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) != 4 {
				continue
			}

			localRef := parts[0]
			remoteSha := parts[3]

			// Is this a new branch being pushed? (remote sha is all zeros)
			isNewBranch := strings.Trim(remoteSha, "0") == ""
			// Is it a branch ref?
			isBranch := strings.HasPrefix(localRef, "refs/heads/")
			// Is it a main/master/develop branch?
			isMainBranch := strings.HasSuffix(localRef, "/main") || strings.HasSuffix(localRef, "/master") || strings.HasSuffix(localRef, "/develop")

			if isBranch && isNewBranch && !isMainBranch {
				fmt.Println("------------------------------------------------")
				fmt.Println("✨ dev-pal detected a new feature branch push!")
				
				// We assume the base branch is 'main' for this agentic feature.
				// This could be made configurable in the future.
				description, err := generatePRDescription("main")
				if err != nil {
					// Don't block the push, just inform the user.
					fmt.Printf("⚠️  Could not auto-generate PR description: %v\n", err)
				} else {
					fmt.Println("✅ Here is a suggested PR description for you:")
					fmt.Println("------------------------------------------------")
					fmt.Println(description)
					fmt.Println("------------------------------------------------")
					fmt.Println("Copy the text above to use in your Pull Request.")
				}
				// We only do this for the first new branch found.
				os.Exit(0)
			}
		}
	},
}

func init() {
	internalPrePushCmd.Flags().StringVar(&remoteName, "remote-name", "", "The name of the remote")
	internalPrePushCmd.Flags().StringVar(&remoteUrl, "remote-url", "", "The URL of the remote")
	rootCmd.AddCommand(internalPrePushCmd)
}
