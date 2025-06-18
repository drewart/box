package cli

import (
	"fmt"
	"os"

	"github.com/drewart/box/internal/config"
	"github.com/drewart/box/internal/util"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [service] [user] [password]",
	Short: "Set a password for a service and user",
	Long:  "Set a password for a service and user using the system keyring",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		user := args[1]
		password := args[2]

		config := config.GetConfig()

		hash := util.HashString(password)

		config.AddUpdateAppUser(service, user, hash, []string{"box"})
		err := config.Save()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving configuration: %s\n", err)
			os.Exit(1)
		}
		// Set password in the keyring
		err = keyring.Set(service, user, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error setting password: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Password for user '%s' in service '%s' set successfully.\n", user, service)
	},
}
