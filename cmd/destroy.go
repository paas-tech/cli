package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "destroy",
	Short:   "Destroy a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Destroying deployment")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
