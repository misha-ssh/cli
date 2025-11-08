package command

import (
	"errors"

	updateForm "github.com/misha-ssh/cli/internal/component/update"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/list"
	"github.com/misha-ssh/cli/internal/component/output"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
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
			return errors.New("not found connections")
		}

		updatedConnection, err := updateForm.Run(selectedConn)
		if err != nil {
			return errors.New("failed to update connections")
		}

		err = kernel.Update(updatedConnection, selectedConn.Alias)
		if err != nil {
			return err
		}

		output.Success("success update connection - ", updatedConnection.Alias)

		return nil
	},
}
