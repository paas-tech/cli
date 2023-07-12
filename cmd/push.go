package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "push",
	Short:   "Push a project repo to PaaSTech",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Pushing code to PaaSTech")

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

		// Open current repo git
		repo, err := git.PlainOpen(".")
		if err != nil {
			return errors.New("No git repository found in current directory")
		}

		remote, _ := repo.Remote("paastech")
		err = GitPush(remote.Config().URLs[0])
		if err != nil {
			return err
		}

		fmt.Println("Project pushed")

		return nil
	},
}

func GitPush(url string) error {
	cmd := exec.Command("git", "push", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
