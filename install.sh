#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Default values
INSTALL_DIR="/usr/local/bin"
REPO_OWNER="wato787"
REPO_NAME="gkit"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="x86_64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo -e "${RED}Unsupported architecture: $ARCH${NC}" && exit 1 ;;
esac

case $OS in
    linux) OS="Linux" ;;
    darwin) OS="Darwin" ;;
    *) echo -e "${RED}Unsupported OS: $OS${NC}" && exit 1 ;;
esac

echo -e "${GREEN}Installing gkit for $OS $ARCH...${NC}"

# Get latest release info
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO_OWNER/$REPO_NAME/releases/latest")
VERSION=$(echo "$LATEST_RELEASE" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
    echo -e "${RED}Failed to get latest release version${NC}"
    exit 1
fi

echo -e "${YELLOW}Latest version: $VERSION${NC}"

# Download URL
DOWNLOAD_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$VERSION/${REPO_NAME}_${OS}_${ARCH}.tar.gz"

# Create temporary directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

echo -e "${YELLOW}Downloading from $DOWNLOAD_URL...${NC}"

# Download and extract
curl -sL "$DOWNLOAD_URL" | tar -xz

# Install binaries
echo -e "${YELLOW}Installing to $INSTALL_DIR...${NC}"

for cmd in gs ga gc gp; do
    if [ -f "$cmd" ]; then
        sudo mv "$cmd" "$INSTALL_DIR/"
        sudo chmod +x "$INSTALL_DIR/$cmd"
        echo -e "${GREEN}✓ Installed $cmd${NC}"
    else
        echo -e "${RED}✗ $cmd not found in archive${NC}"
    fi
done

# Cleanup
cd - > /dev/null
rm -rf "$TMP_DIR"

echo -e "${GREEN}Installation complete!${NC}"
echo -e "${YELLOW}Available commands: gs, ga, gc, gp${NC}"
echo -e "${YELLOW}Run 'gs --help' to get started${NC}"
echo ""
echo -e "${YELLOW}To enable tab completion, add these to your shell config:${NC}"
echo "# For bash (~/.bashrc):"
echo "source <(gs completion bash)"
echo "source <(ga completion bash)"
echo "source <(gc completion bash)"
echo "source <(gp completion bash)"
echo ""
echo "# For zsh (~/.zshrc):"
echo "source <(gs completion zsh)"
echo "source <(ga completion zsh)"
echo "source <(gc completion zsh)"
echo "source <(gp completion zsh)"