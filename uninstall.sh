#!/bin/bash

# ------------------------------------------------------------------
#  UNINSTALLER: Mr.the-ai-dev 😢
# ------------------------------------------------------------------

HOOK_FILE=".git/hooks/post-commit"

echo "😢 Uninstalling Mr.the-ai-dev..."

if [ ! -d ".git" ]; then
    echo "❌ Mamae! Go inside a git folder first."
    exit 1
fi

if [ ! -f "$HOOK_FILE" ]; then
    echo "🤔 Mr.the-ai-dev doesn't seem to be installed here (no post-commit hook found)."
    exit 0 # Not an error, just not installed
fi

# Check if the hook is the one we installed.
if ! grep -q "Mr.the-ai-dev" "$HOOK_FILE"; then
    echo "🤔 The existing post-commit hook doesn't look like Mr.the-ai-dev's."
    echo "Aborting to be safe. Please remove it manually if you are sure."
    exit 1
fi

rm "$HOOK_FILE"

echo "✅ Mr.the-ai-dev has been uninstalled. The post-commit hook is removed."
echo "He'll miss you, Mamae! Come back soon! 🚀"
