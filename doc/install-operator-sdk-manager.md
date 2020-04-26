# Install the Operator SDK CLI

- [Install from GitHub release](#install-from-github-release)
- [Compile and install from master](#compile-and-install-from-master)


## Install from GitHub release

### Download the release binary

```sh
# Set the release version variable
$ RELEASE_VERSION=v0.0.1
# Linux
$ curl -LO https://github.com/akoserwal/operator-sdk-manager/releases/download/${RELEASE_VERSION}/operator-sdk-manager-${RELEASE_VERSION}-x86_64-linux-gnu
# macOS
$ curl -LO https://github.com/akoserwal/operator-sdk-manager/releases/download/${RELEASE_VERSION}/operator-sdk-manager-${RELEASE_VERSION}-x86_64-apple-darwin
```

### Install the release binary in your PATH

```sh
# Linux
$ chmod +x operator-sdk-manager-${RELEASE_VERSION}-x86_64-linux-gnu && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk-manager-${RELEASE_VERSION}-x86_64-linux-gnu /usr/local/bin/operator-sdk-manager && rm operator-sdk-manager-${RELEASE_VERSION}-x86_64-linux-gnu
# macOS
$ chmod +x operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk-manager-${RELEASE_VERSION}-x86_64-apple-darwin /usr/local/bin/operator-sdk-manager && rm operator-sdk-manager-${RELEASE_VERSION}-x86_64-apple-darwin
```

## Compile and install from master

### Prerequisites

- [git][git_tool]
- [go][go_tool] version v1.13+.

```sh
$ git clone https://github.com/akoserwal/operator-sdk-manager
$ cd operator-sdk-manager
$ git checkout master
$ go mod tidy
$ go build
```
[git_tool]:https://git-scm.com/downloads
[go_tool]:https://golang.org/dl/
