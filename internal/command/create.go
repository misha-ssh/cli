package command

import (
	"errors"

	form "github.com/misha-ssh/cli/internal/component/create"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/output"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

// createCmd Command for create create
var createCmd = &cobra.Command{
	Use:   envconst.UseCreateCmd,
	Short: envconst.ShortCreateCmd,
	Long:  envconst.LongCreateCmd,
	RunE: func(cmd *cobra.Command, _ []string) error {
		connection, err := form.Run()
		if err != nil {
			return errors.New("error run text input")
		}

		err = kernel.Create(connection)
		if err != nil {
			return err
		}

		output.Success("success create connection - ", connection.Alias)

		return nil
	},
}
