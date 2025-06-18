package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "box",
	Short: "box cli for ",
	Long:  "box cli for managing secrets and other operations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to box CLI! Use 'box --help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing box '%s'\n", err)
		os.Exit(1)
	}
}
