package main

import (
	"os"

	"github.com/octoproject/octo-cli/command"
	"github.com/spf13/cobra"
)

func main() {
	rootcmd := &cobra.Command{
		Use:   "octo-cli",
		Short: "Expose data from any database as web service",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootcmd.AddCommand(command.NewInitCommand())
	rootcmd.AddCommand(command.NewCreateCommand())
	rootcmd.AddCommand(command.NewBuildCommand())
	rootcmd.AddCommand(command.NewDeployCommand())

	if err := rootcmd.Execute(); err != nil {
		os.Exit(1)
	}
}
