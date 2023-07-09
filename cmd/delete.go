package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/paastech-cloud/cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	GroupID: "project",
	Use:     "delete",
	Short:   "Delete a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deleting this project from PaasTech")

		// confirmation
		stdin := bufio.NewReader(os.Stdin)
		if utils.ConfirmationPrompt(stdin) {
			return project.DeleteProject(viper.GetString("server"), viper.GetString("jwt"), "todo")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
