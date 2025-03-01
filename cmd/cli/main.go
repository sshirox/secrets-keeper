package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "secrets-keeper",
	Short: "CLI secrets keeper",
	Long:  `Secrets keeper for secure storage of logins, passwords and other data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Secret keeper: Use --help for commands.")
	},
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
