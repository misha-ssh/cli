package command

import (
	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/spf13/cobra"
)

// deleteCmd Command for delete connection
var deleteCmd = &cobra.Command{
	Use:   envconst.UseDeleteCmd,
	Short: envconst.ShortDeleteCmd,
	Long:  envconst.LongDeleteCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
