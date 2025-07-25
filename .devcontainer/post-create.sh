#!/bin/bash
set -e



echo "--- Starting Post-Create Command ---"

# Ensure HOME is set correctly for subsequent operations in the container
export HOME="/home/vscode" 

# Install global npm packages (latest versions)
echo "--- Installing npm packages (gemini-cli, pyright)... ---"
npm install -g npm@latest
npm install -g @google/gemini-cli@latest pyright@latest @anthropic-ai/claude-code@latest

# Install Go tools with version logging
echo "--- Installing Go packages (mcp-language-server, gopls)... ---"
echo "Go version: $(go version)"
go install github.com/isaacphi/mcp-language-server@latest
go install golang.org/x/tools/gopls@latest
echo "Installed Go tools to: $(go env GOPATH)/bin"

# Install Rust tools
echo "--- Installing Rust components (rust-analyzer)... ---"
rustup component add rust-analyzer

# Install ElixirLS (latest version with retry logic)
echo "--- Installing ElixirLS... ---"

install_elixir_ls() {
    local max_retries=3
    local retry_count=0
    
    while [ $retry_count -lt $max_retries ]; do
        echo "Attempt $((retry_count + 1)) of $max_retries to install ElixirLS..."
        
        LATEST_ELIXIR_INFO=$(curl -s --connect-timeout 10 --max-time 30 https://api.github.com/repos/elixir-lsp/elixir-ls/releases/latest)
        
        if [ $? -ne 0 ] || [ -z "$LATEST_ELIXIR_INFO" ]; then
            echo "Failed to fetch ElixirLS release info, retrying..."
            retry_count=$((retry_count + 1))
            sleep 5
            continue
        fi
        
        DOWNLOAD_URL=$(echo "$LATEST_ELIXIR_INFO" | jq -r '.assets[] | select(.name | endswith(".zip")) | .browser_download_url')
        
        if [ -z "$DOWNLOAD_URL" ] || [ "$DOWNLOAD_URL" = "null" ]; then
            echo "Failed to find ElixirLS download URL, retrying..."
            retry_count=$((retry_count + 1))
            sleep 5
            continue
        fi
        
        echo "Downloading ElixirLS from: $DOWNLOAD_URL"
        if wget -qO elixir-ls.zip "$DOWNLOAD_URL"; then
            mkdir -p /home/vscode/.local/elixir-ls /home/vscode/.local/bin
            if unzip -o elixir-ls.zip -d "$HOME/.local/elixir-ls"; then
                rm -f elixir-ls.zip
                chmod +x "$HOME/.local/elixir-ls/language_server.sh"
                ln -sf "$HOME/.local/elixir-ls/language_server.sh" "$HOME/.local/bin/elixir-ls"
                echo "ElixirLS installed successfully"
                return 0
            fi
        fi
        
        retry_count=$((retry_count + 1))
        echo "Installation failed, retrying in 5 seconds..."
        sleep 5
    done
    
    echo "ERROR: Failed to install ElixirLS after $max_retries attempts"
    return 1
}

install_elixir_ls || exit 1

# echo "--- Installing ElixirLS... ---"
# mkdir -p "$HOME/.local/elixir-ls"
# mkdir -p "$HOME/.local/bin"
# mkdir -p "$HOME/.cache/elixir-ls"
# ELIXIR_LS_VERSION=$(curl -s https://api.github.com/repos/elixir-lsp/elixir-ls/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
# wget "https://github.com/elixir-lsp/elixir-ls/releases/download/${ELIXIR_LS_VERSION}/elixir-ls-${ELIXIR_LS_VERSION}.zip"
# unzip "elixir-ls-${ELIXIR_LS_VERSION}.zip" -d "$HOME/.local/elixir-ls"
# chmod +x "$HOME/.local/elixir-ls/language_server.sh"
# ln -sf "$HOME/.local/elixir-ls/language_server.sh" "$HOME/.local/bin/elixir-ls"
# rm -f "elixir-ls-${ELIXIR_LS_VERSION}.zip"

# Set ElixirLS environment variables
# export ELS_CACHE_DIR="$HOME/.cache/elixir-ls"
# export ELS_ELIXIR_PATH="/usr/local/bin/elixir"

# Find the pyright-langserver executable and related Node.js paths
echo "Finding pyright-langserver and Node.js paths..."
PYRIGHT_PATH=$(find /usr/local/share/nvm/versions/node -name pyright-langserver | head -n 1)
if [ -z "$PYRIGHT_PATH" ]; then
    echo "ERROR: pyright-langserver not found under /usr/local/share/nvm/versions/node"
    exit 1
fi
echo "Found pyright-langserver at: $PYRIGHT_PATH"

NODE_BIN_DIR=$(dirname "$PYRIGHT_PATH")
NODE_VERSION_DIR=$(dirname "$NODE_BIN_DIR")
NODE_MODULES_PATH="$NODE_VERSION_DIR/lib/node_modules"

# --- Configure Claude Code MCP Servers ---
echo "--- Configuring Claude Code MCP servers using 'claude mcp add'... ---"

# Verify required binaries exist before configuring MCP servers
echo "Verifying language server binaries..."
if [ ! -f "/home/vscode/go/bin/mcp-language-server" ]; then
    echo "ERROR: mcp-language-server not found at /home/vscode/go/bin/mcp-language-server"
    exit 1
fi

if [ ! -f "/home/vscode/go/bin/gopls" ]; then
    echo "ERROR: gopls not found at /home/vscode/go/bin/gopls"
    exit 1
fi

if [ ! -f "$PYRIGHT_PATH" ]; then
    echo "ERROR: pyright-langserver not found at $PYRIGHT_PATH"
    exit 1
fi

if [ ! -f "/home/vscode/.local/bin/elixir-ls" ]; then
    echo "ERROR: elixir-ls not found at /home/vscode/.local/bin/elixir-ls"
    exit 1
fi

if [ ! -f "/home/vscode/.cargo/bin/rust-analyzer" ]; then
    echo "ERROR: rust-analyzer not found at /home/vscode/.cargo/bin/rust-analyzer"
    exit 1
fi

# Create workspace directories if they don't exist
mkdir -p /workspace/dev/go /workspace/dev/python /workspace/dev/elixir /workspace/dev/rust

# Go Server for Claude
echo "Adding Go server MCP for Claude..."
claude mcp add go-server -s project\
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/go \
    --lsp /home/vscode/go/bin/gopls

# Python Server for Claude
echo "Adding Python server MCP for Claude..."
claude mcp add python-server -s project\
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/python \
    --lsp "$PYRIGHT_PATH" \
    -- --stdio

# Elixir Server for Claude
echo "Adding Elixir server MCP for Claude..."
claude mcp add elixir-server -s project\
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/elixir \
    --lsp /home/vscode/.local/bin/elixir-ls

# Rust Server for Claude
echo "Adding Rust server MCP for Claude..."
claude mcp add rust-server -s project\
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/rust \
    --lsp /home/vscode/.cargo/bin/rust-analyzer

# --- Configure Gemini Code Assist MCP Servers ---
echo "--- Configuring Gemini Code Assist MCP servers via /workspace/.gemini/settings.json... ---"

# Ensure /workspace/.gemini/settings.json exists
GEMINI_SETTINGS_DIR="/workspace/.gemini"
GEMINI_SETTINGS_FILE="$GEMINI_SETTINGS_DIR/settings.json"
CLAUDE_MCP_SETTINGS_FILE="/workspace/.mcp.json"

mkdir -p "$GEMINI_SETTINGS_DIR"
if [ ! -f "$GEMINI_SETTINGS_FILE" ]; then
    echo "{}" > "$GEMINI_SETTINGS_FILE"
fi

# Check if .mcp.json exists
if [ ! -f "$CLAUDE_MCP_SETTINGS_FILE" ]; then
    echo "ERROR: /workspace/.mcp.json not found. Cannot configure Gemini MCP servers."
    exit 1
fi

# Use jq to copy mcpServers from .mcp.json to settings.json
jq --slurp '.[1] * {mcpServers: .[0].mcpServers}' "$CLAUDE_MCP_SETTINGS_FILE" "$GEMINI_SETTINGS_FILE" > "$GEMINI_SETTINGS_FILE".tmp && mv "$GEMINI_SETTINGS_FILE".tmp "$GEMINI_SETTINGS_FILE"

# Run version check
echo "--- Running version check... ---"
/workspace/.devcontainer/version-check.sh

echo "--- Dev container is ready! ---"
