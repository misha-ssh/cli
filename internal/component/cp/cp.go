package cp

import (
	"errors"
	"os"

	"github.com/charmbracelet/huh"
)

var (
	errCpFile     = errors.New(`cannot copy file`)
	errGetHomeDir = errors.New(`cannot get home directory`)
)

// Run get form for cp files
func Run() (*Fields, error) {
	var isDownloadCopy bool

	fields := &Fields{}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, errGetHomeDir
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Type copy").
				Description("Choice method copy").
				Affirmative("Download").
				Negative("Upload").
				Value(&isDownloadCopy),
		),
		huh.NewGroup(
			huh.NewFilePicker().
				Title("Local File").
				Description("Select local file").
				CurrentDirectory(homedir).
				Validate(localFileValidate).
				Value(&fields.LocalFile).
				Picking(true),
			huh.NewInput().
				Title("Remote File").
				Description("Absolute path to remote file").
				Value(&fields.RemoteFile),
		).WithHideFunc(func() bool {
			return isDownloadCopy
		}),
		huh.NewGroup(
			huh.NewInput().
				Title("Remote File").
				Description("Absolute path to remote file").
				Value(&fields.RemoteFile),
			huh.NewInput().
				Title("Local File").
				Description("Absolute path to local file").
				Value(&fields.LocalFile),
		).WithHideFunc(func() bool {
			return !isDownloadCopy
		}),
	).WithShowHelp(true).Run()
	if err != nil {
		return nil, errCpFile
	}

	fields.TypeCopy = Upload

	if isDownloadCopy {
		fields.TypeCopy = Download
	}

	return fields, nil
}
