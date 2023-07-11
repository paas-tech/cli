package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "restart",
	Short:   "Restart a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Restarting deployment")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
