# 멀티 언어 개발 환경 - Go 포함
FROM ubuntu:22.04

# 기본 환경 설정
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

# 기본 패키지 설치
RUN apt-get update && apt-get install -y \
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
    && rm -rf /var/lib/apt/lists/*

# Python 설치
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    python3-venv \
    python3-dev \
    && rm -rf /var/lib/apt/lists/*

# Python 심볼릭 링크 생성
RUN ln -s /usr/bin/python3 /usr/local/bin/python

# Go 설치
RUN wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz \
    && rm go1.21.0.linux-amd64.tar.gz

# Go 환경 변수 설정
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GOROOT="/usr/local/go"

# Elixir 설치
RUN apt-get update && apt-get install -y \
    erlang \
    elixir \
    && rm -rf /var/lib/apt/lists/*

# vscode 사용자 생성
RUN groupadd --gid 1000 vscode \
    && useradd --uid 1000 --gid vscode --shell /bin/bash --create-home vscode \
    && echo 'vscode ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# 사용자 전환 후 환경 설정
USER vscode
WORKDIR /home/vscode

# 사용자 환경에서 Go 경로 설정
ENV PATH="/usr/local/go/bin:/home/vscode/go/bin:${PATH}"
ENV GOPATH="/home/vscode/go"

# Rust 설치 (사용자 환경)
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/home/vscode/.cargo/bin:${PATH}"

# 작업 디렉터리 설정
WORKDIR /workspace
USER root
RUN chown -R vscode:vscode /workspace
USER vscode

# 컨테이너 실행 유지
CMD ["sleep", "infinity"]