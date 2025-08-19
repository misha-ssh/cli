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

	updatedConn := *connection

	if len(updatedConn.Password) > 0 {
		authPassConfirm = true
	} else if len(updatedConn.SshOptions.PrivateKey) > 0 {
		authPassConfirm = false
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
			huh.NewFilePicker().
				Title("PrivateKey").
				Description("select file with private key").
				CurrentDirectory(homedir).
				Validate(privateKeyValidate).
				Value(&updatedConn.SshOptions.PrivateKey),
		).WithHideFunc(func() bool {
			return authPassConfirm
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

	if len(connection.Password) > 0 && !authPassConfirm {
		fields.Password = ""
	}

	if len(updatedConn.SshOptions.PrivateKey) > 0 && authPassConfirm {
		fields.PrivateKey = ""
	}

	return fields, nil
}
