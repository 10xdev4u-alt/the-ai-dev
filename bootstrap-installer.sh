#!/bin/bash

set -e # Exit immediately if a command exits with a non-zero status.

# --- Configuration ---
REPO_URL="https://codeberg.org/princetheprogrammerbtw/the-ai-devx.git"
TMP_DIR="/tmp/p2pGit-install"

echo "🚀 Bootstrapping dev-pal installation..."

# --- Dependency Checks ---
echo "Checking for dependencies (git, go)..."
if ! command -v git &> /dev/null; then
    echo "❌ Git is not installed. Please install git to continue."
    exit 1
fi
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to continue."
    exit 1
fi
echo "✅ Dependencies found."

# --- Cloning the Repository ---
echo "Cloning the dev-pal repository to a temporary directory..."
# Remove old temp directory if it exists
if [ -d "$TMP_DIR" ]; then
    rm -rf "$TMP_DIR"
fi
git clone --depth 1 "$REPO_URL" "$TMP_DIR"
cd "$TMP_DIR" || exit 1
echo "✅ Repository cloned."

# --- Running the Local Installer ---
echo "Running the local installer (install-dev-pal.sh)..."
if [ ! -f "./install-dev-pal.sh" ]; then
    echo "❌ The installer script 'install-dev-pal.sh' was not found in the repository."
    exit 1
fi
chmod +x ./install-dev-pal.sh
./install-dev-pal.sh # This script will handle the build and setup

# --- Cleanup ---
echo "Cleaning up temporary files..."
rm -rf "$TMP_DIR"

echo "\n🎉 dev-pal bootstrap installation complete!"
echo "You can now use the 'dev-pal' command globally."
