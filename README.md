# Misha - cli ssh client

![Go Report Card](https://goreportcard.com/badge/github.com/misha-ssh/kernel)
![Release](https://img.shields.io/github/release/misha-ssh/kernel?status.svg)
![Action Lint](https://github.com/misha-ssh/kernel/actions/workflows/lint.yml/badge.svg)
![Action Tests](https://github.com/misha-ssh/kernel/actions/workflows/tests.yml/badge.svg)
![Action Coverage](https://github.com/misha-ssh/kernel/actions/workflows/coverage.yml/badge.svg)
![GitHub top language](https://img.shields.io/github/languages/top/ssh-connection-manager/ssh-)
![GitHub](https://img.shields.io/github/license/ssh-connection-manager/ssh-)
![GitHub Repo stars](https://img.shields.io/github/stars/ssh-connection-manager/ssh-)
![GitHub issues](https://img.shields.io/github/issues/ssh-connection-manager/ssh-)

This package acts as the core for an ssh client written in go

Made using data from packages:

* [crypto](https://pkg.go.dev/golang.org/x/crypto)
* [go-keyring](http://github.com/zalando/go-keyring)
* [term](https://pkg.go.dev/golang.org/x/term)

## ✨ Documentation

install this package in your repository

```bash
go get github.com/misha-ssh/kernel
```

## ✨ Install

install this package in your repository

```bash
go get github.com/misha-ssh/kernel
```

## 📖 Examples & Usage

You will be provided with a list of commands that you can use in your projects

The code with the commands will be on the way - [link](./examples/command)

### 🔌 Connect

The command to connect to the remote server

### ✍️ Create

The command to create a connection

this command saves the connection to a file and goes through the dependency initialization cycle

### 🪄 Update

The command to update the connection

This command also updates the connection data if you need to resave the private key


### 🆑 Delete

The command to delete the connection

This command removes the connection from the file and also deletes the private key if it has been saved

### 📝 List

The command to get a list of connections

This command will list the connections from the previously created connections

## 🧪 Testing

You can run the command for testing after the step with local installation

The command to launch the linter:

```bash
make lint
```

## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)
