package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var internalPreCommitCmd = &cobra.Command{
	Use:    "internal-pre-commit",
	Short:  "Internal command for the pre-commit hook (secret scanner).",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🕵️  dev-pal is scanning for secrets...")

		// 1. Execute git diff
		gitCmd := exec.Command("git", "diff", "--staged", "--diff-filter=d")
		diffOutput, err := gitCmd.Output()
		if err != nil {
			// If git diff fails, we can't check, so we pass but warn the user.
			fmt.Println("⚠️  Warning: Could not run git diff. Skipping secret scan.")
			os.Exit(0)
		}

		// 2. Find suspicious lines
		scanner := bufio.NewScanner(strings.NewReader(string(diffOutput)))
		suspiciousLines := []string{}
		// Regex to find lines being added that contain a suspicious keyword
		re := regexp.MustCompile(`^\+(.*)(secret|key|token|password)`)

		for scanner.Scan() {
			line := scanner.Text()
			if re.MatchString(line) {
				// Add just the content, without the leading '+'
				suspiciousLines = append(suspiciousLines, line[1:])
			}
		}

		// 3. If no suspicious lines, exit cleanly
		if len(suspiciousLines) == 0 {
			fmt.Println("✅ No suspicious keywords found. Good to go.")
			os.Exit(0)
		}

		// 4. If suspicious lines are found, ask the AI
		fmt.Println("🤔 Found some suspicious lines. Asking the AI to verify...")
		prompt := fmt.Sprintf(`You are a security expert. Analyze the following lines of code that were just added to a commit. Determine if any of them contain sensitive credentials, API keys, or secrets. For each line, state whether it is a 'SECRET' or 'SAFE'. Finally, give a single word summary: 'BLOCK' if any line is a secret, and 'PASS' otherwise.

Lines to analyze:
%s`, strings.Join(suspiciousLines, "\n"))

		aiResponse, err := callAI(prompt)
		if err != nil {
			fmt.Printf("⚠️  Warning: AI call failed (%v). Allowing commit to proceed, but please double-check your changes for secrets.\n", err)
			os.Exit(0)
		}

		// 5. Parse the verdict from the AI response
		// A simple approach: check if the last part of the response contains "BLOCK"
		verdict := "PASS"
		if strings.Contains(strings.ToUpper(aiResponse), "BLOCK") {
			verdict = "BLOCK"
		}

		// 6. Act on the verdict
		if verdict == "BLOCK" {
			fmt.Println("------------------------------------------------")
			fmt.Println("🛑 COMMIT REJECTED by dev-pal 🛑")
			fmt.Println("The AI has identified a potential secret in your staged changes:")
			fmt.Println("")
			fmt.Println("--- AI Analysis ---")
			fmt.Println(aiResponse)
			fmt.Println("-------------------")
			fmt.Println("")
			fmt.Println("Please remove the secret from your changes before committing.")
			fmt.Println("To bypass this check for this commit only, use 'git commit --no-verify'.")
			fmt.Println("------------------------------------------------")
			os.Exit(1)
		} else {
			fmt.Println("✅ AI verified the lines and found no secrets. Proceeding.")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(internalPreCommitCmd)
}
