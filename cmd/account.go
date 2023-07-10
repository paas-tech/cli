package cmd

import (
	"fmt"
	"time"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	GroupID: "account",
	Use:     "account",
	Short:   "Get infos about user account",
	RunE: func(cmd *cobra.Command, args []string) error {
		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		jwt, err := config.ExtractJWTInfos(userCfg)
		if err != nil {
			return err
		}

		fmt.Println("👤 You are logged in as: " + jwt.Username)
		fmt.Println("🌐 Server: " + userCfg.GetString("server"))
		timeDiff := jwt.ExpirationTime.Sub(time.Now())
		if timeDiff > 0 {
			fmt.Println(
				"⌛ Current session expires in: " + fmt.Sprintf(
					"%02dh%02dm%02ds",
					int(timeDiff.Hours()),
					int(timeDiff.Minutes())%60,
					int(timeDiff.Seconds())%60,
				),
			)
		} else {
			fmt.Println("❌ Current session is expired. Please log back in.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
