package command

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/version"
	"github.com/spf13/cobra"
)

// Run Start app with cobra cmd
func Run() {
	cmd := &cobra.Command{
		Use:     envconst.UseRootCmd,
		Long:    envconst.LongRootCmd,
		Example: envconst.ExampleRootCmd,
	}

	cmd.Root().CompletionOptions.DisableDefaultCmd = true

	cmd.AddCommand(
		connectCmd,
		createCmd,
		deleteCmd,
		updateCmd,
		cpCmd,
	)

	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithVersion(version.Get()),
	); err != nil {
		os.Exit(1)
	}
}
