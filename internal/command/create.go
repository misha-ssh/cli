package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
)

// createCmd Command for create connection
var createCmd = &cobra.Command{
	Use:   envconst.UseCreateCmd,
	Short: envconst.ShortCreateCmd,
	Long:  envconst.LongCreateCmd,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}
