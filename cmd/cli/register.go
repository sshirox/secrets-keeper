package main

import (
	"fmt"
	"github.com/sshirox/secrets-keeper/internal/cliclient/api"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "New user registration",
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		err := api.Register(email, password)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Successful registration!")
	},
}

func init() {
	mainCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("email", "e", "", "Email")
	registerCmd.Flags().StringP("password", "p", "", "Password")
}
