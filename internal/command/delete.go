package command

import (
	"errors"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/list"
	"github.com/misha-ssh/cli/internal/component/output"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

// deleteCmd Command for delete connection
var deleteCmd = &cobra.Command{
	Use:   envconst.UseDeleteCmd,
	Short: envconst.ShortDeleteCmd,
	Long:  envconst.LongDeleteCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		connections, err := kernel.List()
		if err != nil {
			return err
		}

		selectedConn, err := list.Run(connections)
		if err != nil {
			return errors.New("connection not found")
		}

		err = kernel.Delete(selectedConn)
		if err != nil {
			return err
		}

		output.Success("success delete connection - ", selectedConn.Alias)

		return nil
	},
}
