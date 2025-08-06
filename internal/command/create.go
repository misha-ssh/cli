package command

import (
	"errors"
	"time"

	createForm "github.com/misha-ssh/cli/internal/form/create"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/output"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

var (
	successCreateConnection = "success create connection"

	errRunTextInput = errors.New("error run text input")
)

// createCmd Command for create create
var createCmd = &cobra.Command{
	Use:   envconst.UseCreateCmd,
	Short: envconst.ShortCreateCmd,
	Long:  envconst.LongCreateCmd,
	RunE: func(cmd *cobra.Command, _ []string) error {
		fields, err := createForm.Run()
		if err != nil {
			return errRunTextInput
		}

		connection := &connect.Connect{
			Alias:     fields.Alias,
			Login:     fields.Login,
			Password:  fields.Password,
			CreatedAt: time.Now().Format("2006.01.02 15:04:05"),
			UpdatedAt: time.Now().Format("2006.01.02 15:04:05"),
			Type:      connect.TypeSSH,
			SshOptions: &connect.SshOptions{
				Port:       fields.Port,
				PrivateKey: fields.PrivateKey,
			},
		}

		err = kernel.Create(connection)
		if err != nil {
			return err
		}

		output.Success(successCreateConnection + " - " + connection.Alias)

		return nil
	},
}
