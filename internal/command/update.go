package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
)

// updateCmd Command for update connection
var updateCmd = &cobra.Command{
	Use:   envconst.UseUpdateCmd,
	Short: envconst.ShortUpdateCmd,
	Long:  envconst.LongUpdateCmd,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
