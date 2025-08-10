package command

import (
	"context"
	"errors"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/output"
	"github.com/misha-ssh/cli/internal/version"
	"github.com/spf13/cobra"
)

var (
	errRunApp = errors.New("error running app")
)

// Run Start app with cobra cmd
func Run() {
	cmd := &cobra.Command{
		Use:     envconst.UseRootCmd,
		Long:    envconst.LongRootCmd,
		Example: envconst.ExampleRootCmd,
	}

	// Disable default options cmd
	cmd.Root().CompletionOptions.DisableDefaultCmd = true

	cmd.AddCommand(connectCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(deleteCmd)
	cmd.AddCommand(updateCmd)

	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithVersion(version.Get()),
	); err != nil {
		output.Error(errRunApp.Error())
		os.Exit(1)
	}
}
