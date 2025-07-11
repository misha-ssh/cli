package command

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/cli/configs/envconst"
)

// Run Start app with cobra cmd
func Run() {
	app := &cobra.Command{
		Use:     envconst.UseRootCmd,
		Long:    envconst.LongRootCmd,
		Example: envconst.ExampleRootCmd,
	}

	// Disable default options cmd
	app.Root().CompletionOptions.DisableDefaultCmd = true

	app.AddCommand(connectCmd)
	app.AddCommand(createCmd)
	app.AddCommand(deleteCmd)
	app.AddCommand(listCmd)
	app.AddCommand(updateCmd)

	if err := fang.Execute(context.Background(), app); err != nil {
		os.Exit(1)
	}
}
