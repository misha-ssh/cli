package command

import (
	"time"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/input"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

// createCmd Command for create connection
var createCmd = &cobra.Command{
	Use:   envconst.UseCreateCmd,
	Short: envconst.ShortCreateCmd,
	Long:  envconst.LongCreateCmd,
	RunE: func(cmd *cobra.Command, _ []string) error {
		fields, err := input.Run()
		if err != nil {
			return err
		}

		currentTime := time.Now()

		connection := &connect.Connect{
			Alias:     fields.Alias,
			Login:     fields.Login,
			Password:  fields.Password,
			CreatedAt: currentTime.Format("2006.01.02 15:04:05"),
			UpdatedAt: currentTime.Format("2006.01.02 15:04:05"),
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

		return nil
	},
}
