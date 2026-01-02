#!/bin/bash

echo "🚀 Starting dev-pal installation..."

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "❌ Go is not installed. Please install Go first to build dev-pal."
    echo "   (e.g., visit https://golang.org/doc/install or use your package manager)"
    exit 1
fi

# Navigate to the dev-pal directory
DEV_PAL_DIR="./dev-pal"
if [ ! -d "$DEV_PAL_DIR" ]; then
    echo "❌ dev-pal directory not found at $DEV_PAL_DIR."
    echo "   Please ensure you are running this script from the root of the p2pGit project."
    exit 1
fi
cd "$DEV_PAL_DIR" || exit 1

echo "Building dev-pal binary..."
go build -o dev-pal

if [ $? -ne 0 ]; then
    echo "❌ Failed to build dev-pal. Please check for Go build errors above."
    exit 1
fi

echo "Running dev-pal setup to install it globally..."
echo "This requires administrator privileges to write to /usr/local/bin."
sudo ./dev-pal setup

if [ $? -ne 0 ]; then
    echo "❌ dev-pal setup failed. Please check the output above."
    exit 1
fi

echo "\n🎉 dev-pal is now fully installed and ready to use!"
echo "To activate it in any git repository, navigate to that repo and run: dev-pal install-hooks"
