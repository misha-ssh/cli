package update

import (
	"errors"
	"github.com/misha-ssh/kernel/pkg/connect"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

var (
	errGetHomeDir       = errors.New(`cannot get home directory`)
	errCreateConnection = errors.New(`cannot create connection`)
	errConvertPort      = errors.New(`cannot convert port`)
)

func Run(connection *connect.Connect) (*Fields, error) {
	var authPassConfirm bool

	if len(connection.Password) > 0 {
		authPassConfirm = true
	} else if len(connection.SshOptions.PrivateKey) > 0 {
		authPassConfirm = false
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, errGetHomeDir
	}

	fields := &Fields{}
	port := strconv.Itoa(connection.SshOptions.Port)

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
				Value(&connection.SshOptions.PrivateKey),
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

	fields.Alias = connection.Alias
	fields.Login = connection.Login
	fields.Password = connection.Password
	fields.Port = intPort
	fields.PrivateKey = connection.SshOptions.PrivateKey

	return fields, nil
}
