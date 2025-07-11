package command

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
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
