package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
	"github.com/ssh-connection-manager/cli/internal/component/input"
)

// createCmd Command for create connection
var createCmd = &cobra.Command{
	Use:   envconst.UseCreateCmd,
	Short: envconst.ShortCreateCmd,
	Long:  envconst.LongCreateCmd,
	RunE: func(cmd *cobra.Command, _ []string) error {
		fields, err := input.Init()
		if err != nil {
			return err
		}

		fmt.Println(fields)

		return nil
	},
}
