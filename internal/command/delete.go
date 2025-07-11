package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
)

// deleteCmd Command for delete connection
var deleteCmd = &cobra.Command{
	Use:   envconst.UseDeleteCmd,
	Short: envconst.ShortDeleteCmd,
	Long:  envconst.LongDeleteCmd,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}
