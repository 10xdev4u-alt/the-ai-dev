package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func generatePRDescription(baseBranch string) (string, error) {
	currentBranchBytes, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", fmt.Errorf("could not get current branch: %w", err)
	}
	currentBranch := strings.TrimSpace(string(currentBranchBytes))

	if currentBranch == baseBranch {
		return "", fmt.Errorf("you are on the base branch ('%s'); switch to a feature branch first", baseBranch)
	}

	// Get commit log and diff
	commitLogCmd := exec.Command("git", "log", baseBranch+"..HEAD", "--oneline")
	commitLog, err := commitLogCmd.Output()
	if err != nil || len(commitLog) == 0 {
		return "", fmt.Errorf("could not get commit log or no new commits found on this branch")
	}

	diffStatCmd := exec.Command("git", "diff", baseBranch+"...HEAD", "--stat")
	diffStat, _ := diffStatCmd.Output()

	prompt := fmt.Sprintf(`You are an expert technical writer. Based on the following commit log and diff statistics, generate a Pull Request title and a detailed markdown description.

The title should be a single line summarizing the change.

The description should have two sections:
1. '## What Changed?' - A bulleted list explaining the key changes.
2. '## Why?' - A brief explanation of the reason for these changes.

---
Commit Log:
%s
---
Diff Stat:
%s
---

Now, generate the PR title and description.`, string(commitLog), string(diffStat))

	fmt.Println("🤖 Asking Mr.the-ai-dev to write your PR description...")
	aiResponse, err := callAI(prompt)
	if err != nil {
		return "", fmt.Errorf("AI call failed: %w", err)
	}

	return aiResponse, nil
}
