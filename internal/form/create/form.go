package create

import (
	"os"

	"github.com/charmbracelet/huh"
)

const (
	DefaultPort = "22"
)

func Run() (*Fields, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
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
				Validate(loginValidate).
				Value(&fields.Login),

			huh.NewInput().
				Title("Password").
				EchoMode(huh.EchoModePassword).
				Description("Password to connect to a remote machine").
				Validate(passwordValidate).
				Value(&fields.Password),

			huh.NewInput().
				Title("Port").
				Description("port number to connect to a remote machine").
				Validate(portValidate).
				Value(&fields.Port),

			huh.NewFilePicker().
				Title("PrivateKey").
				Description("select file with private key").
				CurrentDirectory(homedir).
				Validate(privateKeyValidate).
				Value(&fields.PrivateKey),
		),
	).WithShowHelp(true).Run()

	return fields, nil
}
