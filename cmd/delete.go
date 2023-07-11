package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/paastech-cloud/cli/pkg/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	GroupID: "project",
	Use:     "delete",
	Short:   "Delete a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deleting this project from PaasTech")

		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		projCfg, err := config.LoadProjectConfig()
		if err != nil {
			return err
		}

		// confirmation
		stdin := bufio.NewReader(os.Stdin)
		if utils.ConfirmationPrompt(stdin) {
			var project project.Project
			projCfg.UnmarshalKey("project", &project)

			err := project.Delete(userCfg.GetString("server"), userCfg.GetString("jwt"))
			if err != nil {
				return err
			}

			os.Remove("paastech.yaml")
			if err != nil {
				return err
			}

			fmt.Println("Project deleted successfully")
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
