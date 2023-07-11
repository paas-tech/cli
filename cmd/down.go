package cmd

import (
	"errors"
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "down",
	Short:   "Stop a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Stopping deployment")

		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		connected, err := config.IsAuthenticated(userCfg)
		if err != nil {
			return err
		}
		if !connected {
			return errors.New("Not logged in")
		}

		projCfg, err := config.LoadProjectConfig()
		if err != nil {
			return err
		}

		var project project.Project
		projCfg.UnmarshalKey("project", &project)

		err = project.Down(userCfg.GetString("server"), userCfg.GetString("jwt"))
		if err != nil {
			return err
		}

		fmt.Println("Project successfully stopped")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
