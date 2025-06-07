package cli

import (
	"fmt"

	"github.com/drewart/box/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored passwords",
	Long:  "List all stored passwords in the system keyring for the specified service.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {


		cfg, _ := config.Load()

		fmt.Println("App, User, Hashs, Created, Updated, Tags")
		fmt.Println("--------------------------------------------------")
		for _, u := range cfg.AppUserList {
			fmt.Printf("%s, %s, %s, %s, %s, %v\n",u.AppName, u.User, u.PassHash, u.CreaatedAt, u.UpdatedAt, u.Tags)
		}
	},
}