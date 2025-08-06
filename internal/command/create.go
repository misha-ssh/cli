package command

import (
	"errors"
	"strconv"
	"time"

	createForm "github.com/misha-ssh/cli/internal/form/create"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

var (
	errRunTextInput = errors.New("error run text input")
	errConvertPort  = errors.New("error convert port")
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

		currentTime := time.Now()

		port, err := strconv.Atoi(fields.Port)
		if err != nil {
			return errConvertPort
		}

		connection := &connect.Connect{
			Alias:     fields.Alias,
			Login:     fields.Login,
			Password:  fields.Password,
			CreatedAt: currentTime.Format("2006.01.02 15:04:05"),
			UpdatedAt: currentTime.Format("2006.01.02 15:04:05"),
			Type:      connect.TypeSSH,
			SshOptions: &connect.SshOptions{
				Port:       port,
				PrivateKey: fields.PrivateKey,
			},
		}

		err = kernel.Create(connection)
		if err != nil {
			return err
		}

		return nil
	},
}
