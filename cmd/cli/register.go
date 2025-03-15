package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sshirox/secrets-keeper/internal/cliclient/api"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "New user registration",
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		if email == "" || password == "" {
			fmt.Println("Error: Email and password are required.")
			return
		}

		err := api.Register(email, password)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Successful registration!")
	},
}

func init() {
	registerCmd.Flags().StringP("email", "e", "", "Email")
	registerCmd.Flags().StringP("password", "p", "", "Password")
}
