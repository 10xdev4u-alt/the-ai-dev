package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const preCommitHookContent = `#!/bin/sh
# dev-pal: This hook is managed by dev-pal. Do not edit.
dev-pal internal-pre-commit "$@"
`

const postCommitHookContent = `#!/bin/sh
# dev-pal: This hook is managed by dev-pal. Do not edit.
dev-pal internal-post-commit "$@"
`

const prePushHookContent = `#!/bin/sh
# dev-pal: This hook is managed by dev-pal. Do not edit.
# Git passes the remote name and URL as arguments.
# Stdin contains the list of refs being pushed.
dev-pal internal-pre-push --remote-name="$1" --remote-url="$2"
`

var installHooksCmd = &cobra.Command{
	Use:   "install-hooks",
	Short: "Installs git hooks in the current repository.",
	Long:  `Installs the pre-commit, post-commit, and pre-push hooks to the .git/hooks/ directory of the current repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Installing dev-pal git hooks...")

		gitDir, err := findGitDir()
		if err != nil {
			fmt.Println("❌ Error: Not a git repository. Please run this command from within a git project.")
			os.Exit(1)
		}
		hooksDir := filepath.Join(gitDir, "hooks")

		hooks := map[string]string{
			"pre-commit":  preCommitHookContent,
			"post-commit": postCommitHookContent,
			"pre-push":    prePushHookContent,
		}

		for name, content := range hooks {
			hookPath := filepath.Join(hooksDir, name)
			fmt.Printf("Writing %s hook...\n", name)
			// Write the file and set its permissions to be executable
			if err := os.WriteFile(hookPath, []byte(content), 0755); err != nil {
				fmt.Printf("❌ Error: Failed to write %s hook: %v\n", name, err)
				os.Exit(1)
			}
		}

		fmt.Println("\n🎉 dev-pal hooks installed successfully!")
		fmt.Println("The AI assistant is now active in this repository.")
	},
}

// findGitDir searches for the .git directory from the current path upwards.
func findGitDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		gitDir := filepath.Join(dir, ".git")
		info, err := os.Stat(gitDir)
		if err == nil && info.IsDir() {
			return gitDir, nil
		}
		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached the root of the filesystem
			break
		}
		dir = parent
	}
	return "", fmt.Errorf(".git directory not found")
}


func init() {
	rootCmd.AddCommand(installHooksCmd)
}
