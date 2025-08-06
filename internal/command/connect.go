package command

import (
	"errors"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/spf13/cobra"
)

// connectCmd Command for connection by alias connect
var connectCmd = &cobra.Command{
	Use:   envconst.UseConnectCmd,
	Short: envconst.ShortConnectCmd,
	Long:  envconst.LongConnectCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("connect")
	},
}
