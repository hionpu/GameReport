# Multi-language development environment - Including Go
FROM ubuntu:22.04

# Basic environment setup
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

# Install essential utilities first
RUN apt-get update && apt-get install -y \
    bash \
    coreutils \
    util-linux \
    procps \
    curl \
    git \
    build-essential \
    sudo \
    wget \
    unzip \
    software-properties-common \
    apt-transport-https \
    ca-certificates \
    gnupg \
    lsb-release \
    vim \
    nano \
    jq \
    && rm -rf /var/lib/apt/lists/*

# Python installation
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    python3-venv \
    python3-dev \
    && rm -rf /var/lib/apt/lists/*

# Create Python symbolic link
RUN ln -s /usr/bin/python3 /usr/local/bin/python

# Go installation
RUN wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz \
    && rm go1.21.0.linux-amd64.tar.gz

# Go environment variables setup
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GOROOT="/usr/local/go"

# Elixir installation
RUN apt-get update && apt-get install -y \
    erlang \
    elixir \
    && rm -rf /var/lib/apt/lists/*

# Create vscode user with proper setup
RUN groupadd --gid 1000 vscode \
    && useradd --uid 1000 --gid vscode --shell /bin/bash --create-home vscode \
    && echo 'vscode ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Create directories and set permissions
RUN mkdir -p /workspace /home/vscode/go \
    && chown -R vscode:vscode /workspace /home/vscode/go

# Switch to vscode user for remaining setup
USER vscode
WORKDIR /home/vscode

# User-specific environment variables
ENV PATH="/usr/local/go/bin:/home/vscode/go/bin:/home/vscode/.local/bin:${PATH}"
ENV GOPATH="/home/vscode/go"
ENV GOROOT="/usr/local/go"

# Rust installation (user environment)
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/home/vscode/.cargo/bin:${PATH}"

# Set working directory
WORKDIR /workspace

# Use a simpler, more reliable command to keep container running
ENTRYPOINT ["/bin/bash"]
CMD ["-c", "while true; do sleep 10; done"]