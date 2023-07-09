package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	GroupID: "project",
	Use:     "init [Project Name]",
	Short:   "Initialize a project",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Initializing a new project")

		config.LoadAuthConfig()
		project, err := project.CreateProject(viper.GetString("server"), viper.GetString("jwt"), args[0])
		if err != nil {
			return err
		}
		fmt.Println("Project " + project.Name + " created.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
