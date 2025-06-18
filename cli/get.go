package cli

import (
	"fmt"
	"os"

	"github.com/drewart/box/internal/config"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [service] [user]",
	Short: "Get a password for a service and user",
	Long:  "Get a password for a service and user using the system keyring",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		user := args[1]

		cfg, _ := config.Load()

		appUser, _ := cfg.FindAppUser(service, user)

		// Get password from the keyring
		secret, err := keyring.Get(service, user)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting password: %s\n", err)
			os.Exit(1)
		}
		if appUser == nil {
			fmt.Fprintf(os.Stderr, "No app user found for service '%s' and user '%s'\n", service, user)
			os.Exit(1)
		}
		fmt.Printf(`{"app":"%s", "user":"%s", "password":"%s", "created":"%s"}`, service, user, secret, appUser.CreaatedAt)
	},
}
