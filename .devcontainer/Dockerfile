FROM golang:1.22.5-bookworm

WORKDIR /work/app

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# golang-migrate
RUN mkdir golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz > golang-migrate/migrate.tar.gz
RUN tar xvzf golang-migrate/migrate.tar.gz -C golang-migrate/
RUN mv golang-migrate/migrate /usr/bin/migrate
RUN rm -rf golang-migrate

# golangci-lint
# binary will be $(go env GOPATH)/bin/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1

# Allow go pkg installation
RUN chown -R $USERNAME:$USERNAME /go

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update -y \
    && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
    git \
    vim \
    tzdata \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Tokyo

USER $USERNAME

RUN go install github.com/air-verse/air@latest
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/evilmartians/lefthook@latest

# prompt
RUN echo "PS1='\[\e[38;5;45m\]\u\[\e[0m\]@\h:\w \$ '" >> ~/.bashrc
