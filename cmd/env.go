package cmd

import (
	"errors"
	"fmt"

	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	GroupID: "project",
	Use:     "env",
	Short:   "Manage environment variables for project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("not enough arguments")
		}
		return nil
	},
}

var addEnvCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an environment variable to the current projet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return project.AddEnvVar(args[0])
	},
}

var lsEnvCmd = &cobra.Command{
	Use:   "ls",
	Short: "List environment variables of the current projet",
	RunE: func(cmd *cobra.Command, args []string) error {
		vars, err := project.ListEnvVar()
		if err != nil {
			return err
		}
		fmt.Print(vars)
		return nil
	},
}

var rmEnvCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an environment variable from the current projet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return project.RmEnvVar(args[0])
	},
}

func init() {
	envCmd.AddCommand(rmEnvCmd)
	envCmd.AddCommand(lsEnvCmd)
	envCmd.AddCommand(addEnvCmd)
	rootCmd.AddCommand(envCmd)
}
