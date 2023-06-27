package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/deployment"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "status",
	Short:   "List deployments for current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		status, err := deployment.Status()
		if err != nil {
			return err
		}
		fmt.Print(status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
