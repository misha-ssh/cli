package command

import (
	"fmt"
	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/input"
	"github.com/spf13/cobra"
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
