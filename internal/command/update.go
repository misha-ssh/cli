package command

import (
	"errors"
	"strconv"
	"time"

	updateForm "github.com/misha-ssh/cli/internal/form/update"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/form/list"
	"github.com/misha-ssh/cli/internal/output"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

var (
	successUpdateConnection = "success update connection"

	errNotFoundConnection = errors.New("not found connections")
)

// updateCmd Command for update create
var updateCmd = &cobra.Command{
	Use:   envconst.UseUpdateCmd,
	Short: envconst.ShortUpdateCmd,
	Long:  envconst.LongUpdateCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		connections, err := kernel.List()
		if err != nil {
			return err
		}

		selectedConn, err := list.Run(connections)
		if err != nil {
			return err
		}

		for _, conn := range connections.Connects {
			if conn.Alias == selectedConn.Alias {
				fields, err := updateForm.Run(&conn)

				port, err := strconv.Atoi(fields.Port)
				if err != nil {
					return errConvertPort
				}

				updatedConnection := &connect.Connect{
					Alias:     fields.Alias,
					Login:     fields.Login,
					Password:  fields.Password,
					UpdatedAt: time.Now().Format("2006.01.02 15:04:05"),
					Type:      connect.TypeSSH,
					SshOptions: &connect.SshOptions{
						Port:       port,
						PrivateKey: fields.PrivateKey,
					},
				}

				err = kernel.Update(updatedConnection, selectedConn.Alias)
				if err != nil {
					return err
				}

				output.Success(successUpdateConnection + " - " + selectedConn.Alias)

				return nil
			}
		}

		return errNotFoundConnection
	},
}
