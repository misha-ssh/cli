# Misha - cli ssh client

![Go Report Card](https://goreportcard.com/badge/github.com/misha-ssh/cli)
![Release](https://img.shields.io/github/release/misha-ssh/cli?status.svg)
![Action Lint](https://github.com/misha-ssh/cli/actions/workflows/lint.yml/badge.svg)
![Action Build](https://github.com/misha-ssh/cli/actions/workflows/build.yml/badge.svg)
![GitHub Repo stars](https://img.shields.io/github/stars/misha-ssh/cli)

Misha - SSH client made on go

Made using data from packages:

* [cobra](https://github.com/spf13/cobra)
* [fang](http://github.com/charmbracelet/fang)
* [huh](https://github.com/charmbracelet/huh)

## ✨ Documentation

You can read the documentation by clicking on the [link](https://misha-ssh.github.io/docs)

## ✨ Install

Install using homebrew:

```bash
# macOS or Linux
brew install --cask misha
```
You can also install the package from the release via the [link](https://github.com/misha-ssh/cli/releases)

## 📖 Examples & Usage

The list of commands that you can use in this SSH client

### 🔌 Connect

The command to connect to the remote server

[![asciicast](https://asciinema.org/a/734047.svg)](https://asciinema.org/a/734047)

### ✍️ Create

The command to create a connection

[![asciicast](https://asciinema.org/a/734049.svg)](https://asciinema.org/a/734049)

### 🪄 Update

The command to update the connection

[![asciicast](https://asciinema.org/a/734050.svg)](https://asciinema.org/a/734050)

### 🆑 Delete

The command to delete the connection

[![asciicast](https://asciinema.org/a/734051.svg)](https://asciinema.org/a/734051?t=0:05)

### 🤖 Run ssh server

for local testing, you can raise your ssh servers - there are three types of them.

1) password connection

to run, write the command:

```bash
make up-ssh
```

to install and remove the server:

```bash
make down-ssh
```

Server accesses:

* ``login`` - root
* ``address`` - localhost
* ``password`` - password
* ``port`` - 22

2) connect with a private key

to run, write the command:

```bash
make up-ssh-key
```

to install and remove the server:

```bash
make down-ssh-key
```

Server accesses:

* ``login`` - root
* ``address`` - localhost
* ``private key`` - ./dockerkey
* ``port`` - 2222

3) connecting via a non-standard port

to run, write the command:

```bash
make up-ssh-port
```

to install and remove the server:

```bash
make down-ssh-port
```

Server accesses:

* ``login`` - root
* ``address`` - localhost
* ``password`` - password
* ``port`` - 2222


## 🧪 Testing

The command to launch the linter:

```bash
make lint
```

## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)
