package command

import (
	"github.com/octoproject/octo-cli/config"
	"github.com/octoproject/octo-cli/service"
	"github.com/spf13/cobra"
)

type createOptions struct {
	fileName string
}

func NewCreateCommand() *cobra.Command {
	options := createOptions{}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a new service",
		Example: "octo-cli create -f example-config.yml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(options)
		},
		SilenceUsage: true,
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.fileName, "file", "f", "", "Name of octo configuration file")
	return cmd
}
func runCreate(o createOptions) error {
	s, err := config.LoadService(o.fileName)
	if err != nil {
		return err
	}

	err = service.NewFunction(s)
	if err != nil {
		return err
	}
	return nil
}
