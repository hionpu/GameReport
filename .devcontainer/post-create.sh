#!/bin/bash
set -e

echo "--- Starting Post-Create Command ---"

# Ensure HOME is set correctly for subsequent operations in the container
export HOME="/home/vscode" 

# Install global npm packages
echo "--- Installing npm packages (gemini-cli, pyright)... ---"
npm install -g @google/gemini-cli pyright @anthropic-ai/claude-code

# Install Go tools
echo "--- Installing Go packages (mcp-language-server, gopls)... ---"
go install github.com/isaacphi/mcp-language-server@latest
go install golang.org/x/tools/gopls@latest

# Install Rust tools
echo "--- Installing Rust components (rust-analyzer)... ---"
rustup component add rust-analyzer

# Install ElixirLS
echo "--- Installing ElixirLS... ---"
mkdir -p "$HOME/.local/elixir-ls"
mkdir -p "$HOME/.local/bin"
wget https://github.com/elixir-lsp/elixir-ls/releases/download/v0.21.0/elixir-ls-v0.21.0.zip
unzip elixir-ls-v0.21.0.zip -d "$HOME/.local/elixir-ls"
chmod +x "$HOME/.local/elixir-ls/language_server.sh"
ln -sf "$HOME/.local/elixir-ls/language_server.sh" "$HOME/.local/bin/elixir-ls"
rm -f elixir-ls-v0.21.0.zip

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
mkdir -p /workspace/dev/go /workspace/dev/python /workspace/dev/elixir_test /workspace/dev/rust

# Go Server for Claude
echo "Adding Go server MCP for Claude..."
claude mcp add go-server \
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/go \
    --lsp /home/vscode/go/bin/gopls

# Python Server for Claude
echo "Adding Python server MCP for Claude..."
claude mcp add python-server \
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/python \
    --lsp "$PYRIGHT_PATH" \
    -- --stdio

# Elixir Server for Claude
echo "Adding Elixir server MCP for Claude..."
claude mcp add elixir-server \
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/elixir_test \
    --lsp /home/vscode/.local/bin/elixir-ls

# Rust Server for Claude
echo "Adding Rust server MCP for Claude..."
claude mcp add rust-server \
    -- /home/vscode/go/bin/mcp-language-server \
    --workspace /workspace/dev/rust \
    --lsp /home/vscode/.cargo/bin/rust-analyzer

# --- Configure Gemini Code Assist MCP Servers ---
echo "--- Configuring Gemini Code Assist MCP servers via ~/.gemini/settings.json... ---"

# Ensure ~/.gemini/settings.json exists
SETTINGS_DIR="$HOME/.gemini"
SETTINGS_FILE="$SETTINGS_DIR/settings.json"
mkdir -p "$SETTINGS_DIR"
if [ ! -f "$SETTINGS_FILE" ]; then
    echo "{}" > "$SETTINGS_FILE"
fi

# Use jq to add/update the mcpServers key in ~/.gemini/settings.json
jq --arg pyright_path "$PYRIGHT_PATH" --arg node_bin_dir "$NODE_BIN_DIR" --arg node_modules_path "$NODE_MODULES_PATH" '.mcpServers =
{
    "go-server": {
      "command": "/home/vscode/go/bin/mcp-language-server",
      "args": ["--workspace", "/workspace/dev/go", "--lsp", "gopls"],
      "env": {"PATH": "/home/vscode/go/bin:/usr/local/go/bin:/usr/bin:/bin", "GOPATH": "/home/vscode/go", "GOCACHE": "/home/vscode/.cache/go-build", "GOMODCACHE": "/home/vscode/go/pkg/mod"}
    },
    "python-server": {
      "command": "/home/vscode/go/bin/mcp-language-server",
      "args": ["--workspace", "/workspace/dev/python", "--lsp", $pyright_path, "--", "--stdio"],
      "env": {"PATH": ($node_bin_dir + ":/usr/bin:/bin:/usr/local/go/bin"), "NODE_PATH": $node_modules_path}
    },
    "elixir-server": {
        "command": "/home/vscode/go/bin/mcp-language-server",
        "args": ["--workspace", "/workspace/dev/elixir_test", "--lsp", "/home/vscode/.local/bin/elixir-ls"],
        "env": {"PATH": "/home/vscode/.local/bin:/usr/bin:/bin"}
    },
    "rust-server": {
        "command": "/home/vscode/go/bin/mcp-language-server",
        "args": ["--workspace", "/workspace/dev/rust", "--lsp", "rust-analyzer"],
        "env": {"PATH": "$HOME/.cargo/bin:$PATH"}
    }
}' "$SETTINGS_FILE" > "$SETTINGS_FILE".tmp && mv "$SETTINGS_FILE".tmp "$SETTINGS_FILE"

echo "--- Dev container is ready! ---"
