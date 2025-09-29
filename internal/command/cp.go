package command

import (
	"fmt"

	"github.com/misha-ssh/cli/configs/envconst"
	"github.com/misha-ssh/cli/internal/component/cp"
	"github.com/misha-ssh/cli/internal/component/list"
	"github.com/misha-ssh/cli/internal/component/output"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/spf13/cobra"
)

// cpCmd Command for copy file
var cpCmd = &cobra.Command{
	Use:   envconst.UseCpCmd,
	Short: envconst.ShortCpCmd,
	Long:  envconst.LongCpCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		connections, err := kernel.List()
		if err != nil {
			return err
		}

		selectedConn, err := list.Run(connections)
		if err != nil {
			return errNotFoundConnection
		}

		fields, err := cp.Run()
		if err != nil {
			return err
		}

		switch fields.TypeCopy {
		case cp.Download:
			err := kernel.Download(selectedConn, fields.RemoteFile, fields.LocalFile)
			if err != nil {
				return err
			}

			output.Success("Success download file: " + fields.LocalFile)
		case cp.Upload:
			err := kernel.Upload(selectedConn, fields.LocalFile, fields.RemoteFile)
			if err != nil {
				return err
			}

			output.Success("Success upload file: " + fields.LocalFile)
		default:
			return fmt.Errorf("unknown copy type: %s", fields.TypeCopy)
		}

		return nil
	},
}
