package cmd

import (
	"errors"
	"strconv"

	"github.com/paastech-cloud/cli/pkg/deployment"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "scale [number]",
	Short:   "Scale a project deployment",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		replicas, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("not a correct number")
		}
		return deployment.Scale(replicas)
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
}
