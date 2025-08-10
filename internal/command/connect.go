package command

import (
	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/list"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

// connectCmd Command for connection by alias connect
var connectCmd = &cobra.Command{
	Use:   envconst.UseConnectCmd,
	Short: envconst.ShortConnectCmd,
	Long:  envconst.LongConnectCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		connections, err := kernel.List()
		if err != nil {
			return err
		}

		selectedConn, err := list.Run(connections)
		if err != nil {
			return err
		}

		err = kernel.Connect(selectedConn)
		if err != nil {
			return err
		}

		return nil
	},
}
