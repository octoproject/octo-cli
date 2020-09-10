package command

import (
	"errors"

	"github.com/octoproject/octo-cli/config"
	"github.com/octoproject/octo-cli/service"
	"github.com/spf13/cobra"
)

var (
	ErrEmptyRegistryPrefix = errors.New("error empty Docker Registry prefix")
)

type buildOptions struct {
	fileName       string
	registryPrefix string
	imageTag       string
}

func NewBuildCommand() *cobra.Command {
	options := buildOptions{}

	cmd := &cobra.Command{
		Use:     "build",
		Short:   "Build function Docker container",
		Example: "octo-cli build --file example-config.yml --prefix test.com --tag v1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBuild(options)
		},
		SilenceUsage: true,
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.fileName, "file", "f", "", "Name of octo configuration file")
	flags.StringVarP(&options.registryPrefix, "prefix", "p", "", "Docker Registry prefix to build image")
	flags.StringVarP(&options.imageTag, "tag", "t", "", "Name and optionally a tag in the 'name:tag' format")

	return cmd
}
func runBuild(o buildOptions) error {
	s, err := config.LoadService(o.fileName)
	if err != nil {
		return err
	}

	if len(o.registryPrefix) < 1 {
		return ErrEmptyRegistryPrefix
	}
	err = service.BuildFunction(s, o.registryPrefix, o.imageTag)
	if err != nil {
		return err
	}
	return nil
}
