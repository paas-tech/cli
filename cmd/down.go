package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "down",
	Short:   "Stop a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Stopping deployment")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
