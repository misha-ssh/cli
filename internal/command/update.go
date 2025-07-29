package command

import (
	"fmt"
	"github.com/misha-ssh/cli/internal/form/list"
	"github.com/misha-ssh/kernel/pkg/kernel"

	"github.com/misha-ssh/cli/configs/envconst"
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

		selectedConnection, err := list.Run(connections)
		if err != nil {
			return err
		}

		fmt.Println(selectedConnection)

		return nil
	},
}
