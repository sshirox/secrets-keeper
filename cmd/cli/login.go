package main

import (
	"fmt"
	"github.com/sshirox/secrets-keeper/config"
	"github.com/sshirox/secrets-keeper/internal/cliclient/api"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User login",
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		token, err := api.Login(email, password)
		if err != nil {
			fmt.Println("Login error:", err)
			return
		}

		config.SaveToken(token)
		fmt.Println("Login successful! Token saved.")
	},
}

func init() {
	loginCmd.Flags().StringP("email", "e", "", "Email")
	loginCmd.Flags().StringP("password", "p", "", "Password")
}
