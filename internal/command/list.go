package command

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
)

// listCmd Command for get list connections
var listCmd = &cobra.Command{
	Use:   envconst.UseListCmd,
	Short: envconst.ShortListCmd,
	Long:  envconst.LongListCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("list")
	},
}
