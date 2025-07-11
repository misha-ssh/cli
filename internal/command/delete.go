package command

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
)

// deleteCmd Command for delete connection
var deleteCmd = &cobra.Command{
	Use:   envconst.UseDeleteCmd,
	Short: envconst.ShortDeleteCmd,
	Long:  envconst.LongDeleteCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("delete")
	},
}
