package update

import (
	"errors"
	"github.com/misha-ssh/kernel/pkg/connect"
	"os"

	"github.com/charmbracelet/huh"
)

const (
	DefaultPort = "22"
)

var (
	errGetHomeDir       = errors.New(`cannot get home directory`)
	errCreateConnection = errors.New(`cannot create connection`)
)

func Run(connection *connect.Connect) (*Fields, error) {
	var authPassConfirm bool

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, errGetHomeDir
	}

	fields := &Fields{
		Port: DefaultPort,
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Alias").
				Description("Unique connection name").
				Validate(aliasValidate).
				Value(&fields.Alias),

			huh.NewInput().
				Title("Login").
				Description("Username of the remote machine").
				Validate(huh.ValidateNotEmpty()).
				Value(&fields.Login),

			huh.NewInput().
				Title("Port").
				Description("Port number to connect to a remote machine").
				Validate(portValidate).
				Value(&fields.Port),

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
				Value(&fields.Password),
		).WithHideFunc(func() bool {
			return !authPassConfirm
		}),
		huh.NewGroup(
			huh.NewFilePicker().
				Title("PrivateKey").
				Description("select file with private key").
				CurrentDirectory(homedir).
				Validate(privateKeyValidate).
				Value(&fields.PrivateKey),
		).WithHideFunc(func() bool {
			return authPassConfirm
		}),
	).WithShowHelp(true).Run()
	if err != nil {
		return nil, errCreateConnection
	}

	return fields, nil
}
