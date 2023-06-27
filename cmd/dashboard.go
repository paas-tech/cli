package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open dashboard on web UI",
	RunE: func(cmd *cobra.Command, args []string) error {
		// only support linux for now
		return exec.Command("xdg-open", "https://paastech.cloud/#/dashboard").Start()
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
