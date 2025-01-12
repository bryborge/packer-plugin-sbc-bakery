FROM golang:1.23.4
ARG USERNAME=vscode

# Setup user
RUN adduser $USERNAME && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# Install packages and Go language server
ARG PKGS="\
git \
goreleaser \
openssh-client \
vim \
zsh \
"

RUN echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list && \
    apt update && \
    apt install --no-install-recommends -y ${PKGS} && \
    apt upgrade -y

WORKDIR /app
