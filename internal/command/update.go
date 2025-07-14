package command

import (
	"errors"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/spf13/cobra"
)

// updateCmd Command for update connection
var updateCmd = &cobra.Command{
	Use:   envconst.UseUpdateCmd,
	Short: envconst.ShortUpdateCmd,
	Long:  envconst.LongUpdateCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("update")
	},
}
