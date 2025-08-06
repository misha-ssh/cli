package command

import (
	"errors"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/spf13/cobra"
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
