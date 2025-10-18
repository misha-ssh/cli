<p align="center">
  <img width="180" alt="mascot" src=".assets/mascot.svg">
</p>

# Misha - tui ssh client

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

## ✨ Install

Install using homebrew:

```bash
# macOS or Linux
brew install misha-ssh/tap/misha
```

>If you get an error - ``git-credential-osxkeychain wants to access key``
then you can refuse to provide access and continue downloading

You can also install the package from the release via the [link](https://github.com/misha-ssh/cli/releases)

## 📖 Examples & Usage

The list of commands that you can use in this SSH client

### 🔌 Connect

The command to connect to the remote server

[![video](.assets/connect.svg)](https://asciinema.org/a/734047)

### 📥 Copy

The command to upload | download files

[![video](.assets/copy.svg)](https://asciinema.org/a/745884)

### ✍️ Create

The command to create a connection

[![video](.assets/create.svg)](https://asciinema.org/a/734430)

### 🪄 Update

The command to update the connection

[![video](.assets/update.svg)](https://asciinema.org/a/734431)

### 🆑 Delete

The command to delete the connection

[![video](.assets/delete.svg)](https://asciinema.org/a/734051)

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

## ⌨️ Local Run

The command to build app:

```bash
make build
```

Run bin file for run app:

```bash
./misha
```

## 🧪 Testing

The command to launch the linter:

```bash
make lint
```

## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)
