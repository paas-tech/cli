package cmd

import (
	"errors"

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
		return nil
	},
}

var lsEnvCmd = &cobra.Command{
	Use:   "ls",
	Short: "List environment variables of the current projet",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var rmEnvCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an environment variable from the current projet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	envCmd.AddCommand(rmEnvCmd)
	envCmd.AddCommand(lsEnvCmd)
	envCmd.AddCommand(addEnvCmd)
	rootCmd.AddCommand(envCmd)
}
