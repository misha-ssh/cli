package update

import (
	"errors"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/misha-ssh/kernel/pkg/connect"
)

var (
	errGetHomeDir       = errors.New(`cannot get home directory`)
	errCreateConnection = errors.New(`cannot create connection`)
	errConvertPort      = errors.New(`cannot convert port`)
)

// Run get form for update connect.Connect
func Run(connection *connect.Connect) (*Fields, error) {
	var authPassConfirm bool
	var updatedPrivateKey bool

	updatedConn := *connection

	hasOriginalPrivateKey := len(connection.SshOptions.PrivateKey) > 0
	hasOriginalPassword := len(connection.Password) > 0

	if hasOriginalPassword {
		authPassConfirm = true
	} else if hasOriginalPrivateKey {
		authPassConfirm = false
	} else {
		authPassConfirm = true
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, errGetHomeDir
	}

	fields := &Fields{}
	port := strconv.Itoa(updatedConn.SshOptions.Port)

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Alias").
				Description("Unique updatedConn name").
				Validate(func(newAlias string) error {
					if connection.Alias != newAlias {
						return aliasValidate(newAlias)
					}

					return nil
				}).
				Value(&updatedConn.Alias),

			huh.NewInput().
				Title("Login").
				Description("Username of the remote machine").
				Validate(huh.ValidateNotEmpty()).
				Value(&updatedConn.Login),

			huh.NewInput().
				Title("Address").
				Description("Address of the remote machine").
				Validate(huh.ValidateNotEmpty()).
				Value(&updatedConn.Address),

			huh.NewInput().
				Title("Port").
				Description("Port number to connect to a remote machine").
				Validate(portValidate).
				Value(&port),

			huh.NewConfirm().
				Title("Authentication").
				Description("Choice method authentication").
				Affirmative("Password").
				Negative("Private key").
				Value(&authPassConfirm),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Password").
				EchoMode(huh.EchoModePassword).
				Description("Password to connect to a remote machine").
				Value(&updatedConn.Password),
		).WithHideFunc(func() bool {
			return !authPassConfirm
		}),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Update private key?").
				Description("Do you want to update the private key?").
				Affirmative("Yes").
				Negative("No").
				Value(&updatedPrivateKey),
		).WithHideFunc(func() bool {
			return authPassConfirm || !hasOriginalPrivateKey
		}),
		huh.NewGroup(
			huh.NewFilePicker().
				Title("PrivateKey").
				Description("Select file with private key").
				CurrentDirectory(homedir).
				Validate(privateKeyValidate).
				Value(&updatedConn.SshOptions.PrivateKey).
				Picking(true),
		).WithHideFunc(func() bool {
			if authPassConfirm {
				return true
			}

			if hasOriginalPrivateKey {
				return !updatedPrivateKey
			}

			return false
		}),
	).WithShowHelp(true).Run()
	if err != nil {
		return nil, errCreateConnection
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, errConvertPort
	}

	fields.Alias = updatedConn.Alias
	fields.Login = updatedConn.Login
	fields.Address = updatedConn.Address
	fields.Port = intPort
	fields.Password = updatedConn.Password
	fields.PrivateKey = updatedConn.SshOptions.PrivateKey

	if hasOriginalPrivateKey && !updatedPrivateKey && !authPassConfirm {
		fields.PrivateKey = connection.SshOptions.PrivateKey
	}

	if hasOriginalPassword && !authPassConfirm {
		fields.Password = ""
	}

	if hasOriginalPrivateKey && authPassConfirm {
		fields.PrivateKey = ""
	}

	return fields, nil
}
