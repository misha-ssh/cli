# Misha - cli ssh client

[![Go Report Card](https://goreportcard.com/badge/github.com/misha-ssh/kernel)](https://goreportcard.com/report/github.com/misha-ssh/kernel)
[![Go Docs](https://godoc.org/github.com/misha-ssh/kernel?status.svg)](https://godoc.org/github.com/misha-ssh/kernel)
[![Release](https://img.shields.io/github/release/misha-ssh/kernel?status.svg)](https://github.com/misha-ssh/kernel/releases)
[![Action Lint](https://github.com/misha-ssh/kernel/actions/workflows/lint.yml/badge.svg)](https://github.com/misha-ssh/kernel)
[![Action Tests](https://github.com/misha-ssh/kernel/actions/workflows/tests.yml/badge.svg)](https://github.com/misha-ssh/kernel)
[![Action Coverage](https://github.com/misha-ssh/kernel/actions/workflows/coverage.yml/badge.svg)](https://github.com/misha-ssh/kernel)

This package acts as the core for an ssh client written in go

Made using data from packages:

* [crypto](https://pkg.go.dev/golang.org/x/crypto)
* [go-keyring](http://github.com/zalando/go-keyring)
* [term](https://pkg.go.dev/golang.org/x/term)

## 📝 Features

- **Connection Management:** Commands for creating, connecting, deleting, and updating your connection
- **Data encryption:** Your connection is securely encrypted
- **Configurations:** Possibility of connection configuration
- **The local environment** There is an environment for testing the connection
- **Flexibility** The ability to embed in any client on go

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

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Connect(connection)
	if err != nil {
		panic(err)
	}
}
```

### ✍️ Create

The command to create a connection

this command saves the connection to a file and goes through the dependency initialization cycle

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Create(connection)
	if err != nil {
		panic(err)
	}
}
```

### 🪄 Update

The command to update the connection

This command also updates the connection data if you need to resave the private key

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Update(connection, "test")
	if err != nil {
		panic(err)
	}
}
```

### 🆑 Delete

The command to delete the connection

This command removes the connection from the file and also deletes the private key if it has been saved

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Delete(connection)
	if err != nil {
		panic(err)
	}
}
```

### 📝 List

The command to get a list of connections

This command will list the connections from the previously created connections

```go
package main

import (
	"fmt"

	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connections, err := kernel.List()
	if err != nil {
		panic(err)
	}

	fmt.Println(connections)
}
```

### 🖥 Struct Connection

This structure describes our connection

We pass this structure to the commands:

- Connect
- Create
- Update
- Delete

Description of fields:

* ``Alias`` - unique name (we use it when selecting an exception from the list and create unique connections to identify
  them)
* ``Login`` - the user's login on the remote device
* ``Password`` - if you have a password connection, then fill in this field, if with a key, then leave an empty line.
* ``Address`` - the address of the remote device
* ``Type`` - there is only one connection type so far - ``connect.TypeSSH``
* ``CreatedAt`` - the creation time is filled in manually
* ``UpdatedAt`` - the update time is filled in manually
* ``SshOptions`` - this is a structure with additional fields for creating software - ``connect.TypeSSH``
* ``Port`` - the port is filled in manually
* ``PrivateKey`` - the path along with the name of the private key

```go
package main

import (
	"fmt"
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
)

func main() {
	connection := &connect.Connect{
		Alias:     "test",
		Login:     "root",
		Password:  "password",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: time.Now().Format("2006.01.02 15:04:05"),
		UpdatedAt: time.Now().Format("2006.01.02 15:04:05"),
		SshOptions: &connect.SshOptions{
			Port:       22,
			PrivateKey: "/path/to/private/key",
		},
	}

	fmt.Println(connection)
}
```

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

### 🔖 Description variable

The variables that the application uses are located here:

* App values - [link](configs/envconst)
* Config keys - [link](configs/envname)

App values (the values that are used in the application and also in the config):

* ``AppName`` - project name & project directory
* ``Theme`` - The theme is an application, there is no implementation at this stage.
* ``DirectionPrivateKeys`` - the name of the directory where the keys will be saved
* ``FilenameConnections`` - the name of the connection file
* ``FilenameConfig`` - the name of the file with the application configs
* ``NameServiceCryptKey`` - the names of the service that will store the private key for encryption
* ``TypeConsoleLogger`` - type for console logging
* ``TypeStorageLogger`` - the type for logging to a file
* ``TypeCombinedLogger`` - type for all types of logging

Config keys (these keys are located in the application configuration):

* ``Theme`` - stores the application's theme, in this case there is no implementation
* ``Logger`` - stores the type of application logging

## 🧪 Testing

You can run the command for testing after the step with local installation

The command to launch the linter:

```bash
make lint
```

Run Unit tests:

```bash
make tests
```

Run test coverage:

```bash
make test-coverage
```

## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)
