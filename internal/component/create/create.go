package create

import (
	"errors"
	"github.com/misha-ssh/kernel/pkg/connect"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

const (
	DefaultPort = "22"
)

var (
	errGetHomeDir       = errors.New(`cannot get home directory`)
	errCreateConnection = errors.New(`cannot create connection`)
	errConvertPort      = errors.New(`cannot convert port`)
)

// Run Main function that runs an interactive form to collect SSH connection details
func Run() (*connect.Connect, error) {
	var authPassConfirm bool

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, errGetHomeDir
	}

	connection := new(connect.Connect)

	port := DefaultPort

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Alias").
				Description("Unique connection name").
				Validate(aliasValidate).
				Value(&connection.Alias),

			huh.NewInput().
				Title("Login").
				Description("Username of the remote machine").
				Validate(huh.ValidateNotEmpty()).
				Value(&connection.Login),

			huh.NewInput().
				Title("Address").
				Description("Address of the remote machine").
				Validate(huh.ValidateNotEmpty()).
				Value(&connection.Address),

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
				Value(&connection.Password),
		).WithHideFunc(func() bool {
			return !authPassConfirm
		}),
		huh.NewGroup(
			huh.NewFilePicker().
				Title("PrivateKey").
				Description("select file with private key").
				CurrentDirectory(homedir).
				Validate(privateKeyValidate).
				Value(&connection.SshOptions.PrivateKey).
				Picking(true),
			huh.NewInput().
				Title("Passphrase").
				EchoMode(huh.EchoModePassword).
				Description("Passphrase for private key (may be empty)").
				Value(&connection.SshOptions.Passphrase),
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

	connection.SshOptions.Port = intPort

	return connection, nil
}
