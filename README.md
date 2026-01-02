# dev-pal PAL 🤖

**Your friendly, AI-powered development assistant.**

---

`dev-pal` is a command-line tool that acts as your "pair-programming" assistant, using AI to automate tedious development tasks and provide intelligent feedback right inside your git workflow. It's designed to be your sidekick, helping you maintain high-quality code and a clean repository with minimal effort.

## ✨ Features

- **AI-Powered Secret Scanning:** Automatically scans your staged changes for API keys, passwords, and other secrets before you commit. It uses AI to verify findings and prevents you from accidentally leaking credentials.
- **Automated Tech-Debt Tracking:** Finds `TODO:`, `FIXME:`, or `HACK:` comments in your commits and automatically logs them in a `TECH_DEBT.md` file, ensuring they never get lost.
- **Agentic PR Descriptions:** When you push a new feature branch, `dev-pal` automatically generates a detailed Pull Request title and description for you to use.
- **AI-Powered Branch Naming:** Use the `new-feature` command to generate clean, conventional branch names from a simple description.
- **Motivational "Vibes":** Get a personalized, supportive message from your AI pal after every commit to keep the energy high.

## 🚀 Installation

`dev-pal` is a self-contained Go application. Installation is handled by a single script.

**Prerequisites:**
- You must have [Go](https://golang.org/doc/install) (version 1.18 or newer) installed.
- You must be in a `bash` or `zsh` compatible shell.

**Installation Steps:**

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/MrPrinceTheProgrammer/p2pGit.git
    cd p2pGit
    ```

2.  **Run the installer:**
    This script will build the `dev-pal` binary and install it globally on your system. It may prompt for administrator (`sudo`) privileges to copy the binary to `/usr/local/bin`.
    ```bash
    ./install-dev-pal.sh
    ```

## 🛠️ Usage

Once installed, `dev-pal` is ready to be activated in any of your Git repositories.

### 1. Activate in a Repository

Navigate to any Git repository where you want to use the assistant and run:
```bash
dev-pal install-hooks
```
This will install the `pre-commit`, `post-commit`, and `pre-push` hooks for that specific repository. The assistant is now active!

### 2. Use the Commands

#### Create a new feature branch:
```bash
dev-pal new-feature "add user authentication with google"
# ✅ AI suggested branch name: feat/user-auth-google
# Creating and switching to the new branch...
# 🚀 Done! You are now on branch 'feat/user-auth-google'.
```

#### Manually generate a PR description:
(The agentic pre-push hook does this automatically, but you can trigger it manually if needed.)
```bash
# Make sure you are on your feature branch
dev-pal generate-pr-desc --base main
```

### 3. The Agentic Workflow

Once the hooks are installed in your repo, the magic happens automatically:
- **`git commit`:** The secret scanner will run before your commit is saved. After it's saved, the tech-debt tracker and vibe generator will run.
- **`git push`:** When you push a new branch for the first time, the PR description generator will run.

## 📜 Legacy Scripts

This repository also contains the original shell scripts (`install.sh`, `uninstall.sh`, etc.) that were the first version of this project. They are preserved here for historical reference but are superseded by the more powerful `dev-pal` Go application.

## 📄 License

This project is licensed under the MIT License.
