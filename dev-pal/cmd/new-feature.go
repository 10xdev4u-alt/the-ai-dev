package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var newFeatureCmd = &cobra.Command{
	Use:   "new-feature [description]",
	Short: "Creates a new feature branch with an AI-generated name.",
	Long:  `Takes a feature description, asks the AI to generate a suitable branch name, and then creates and checks out the new branch.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		fmt.Println("🧠 Asking Mr.the-ai-dev to come up with a cool branch name...")

		prompt := fmt.Sprintf(`You are an expert software engineer. Based on the following feature description, generate a concise, lowercase, kebab-case git branch name. The name should start with 'feat/', 'fix/', or 'chore/'.

Description: "%s"

Branch name:`, description)

		aiResponse, err := callAI(prompt)
		if err != nil {
			fmt.Printf("⚠️  Mr.the-ai-dev is offline. Could not generate a branch name. Error: %v\n", err)
			os.Exit(1)
		}

		// Sanitize the branch name from the AI response
		branchName := sanitizeBranchName(aiResponse)

		if branchName == "" {
			fmt.Println("⚠️  The AI returned an empty branch name. Please try again.")
			os.Exit(1)
		}

		fmt.Printf("✅ AI suggested branch name: %s\n", branchName)
		fmt.Println("Creating and switching to the new branch...")

		gitCmd := exec.Command("git", "checkout", "-b", branchName)
		output, err := gitCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Failed to create branch. Error: %v\n", err)
			fmt.Println(string(output))
			os.Exit(1)
		}

		fmt.Printf("🚀 Done! You are now on branch '%s'. Happy coding, Mamae!\n", branchName)
	},
}

func sanitizeBranchName(input string) string {
	// Get the last line of the response
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sanitized := lines[len(lines)-1]
	
	// Convert to lowercase
	sanitized = strings.ToLower(sanitized)
	
	// Replace spaces and underscores with hyphens
	re := regexp.MustCompile(`[ _]+`)
	sanitized = re.ReplaceAllString(sanitized, "-")
	
	// Remove any characters that are not alphanumeric, hyphen, or slash
	re = regexp.MustCompile(`[^a-z0-9\/-]`)
	sanitized = re.ReplaceAllString(sanitized, "")

	return sanitized
}

func init() {
	rootCmd.AddCommand(newFeatureCmd)
}

