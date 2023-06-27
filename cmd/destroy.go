package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/deployment"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "destroy",
	Short:   "Destroy a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Destroying deployment")
		return deployment.Destroy()
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
