package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sshirox/secrets-keeper/internal/cliclient/api"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition new secrets to the secret store",
	Run: func(cmd *cobra.Command, args []string) {
		secretType, _ := cmd.Flags().GetString("type")
		data, _ := cmd.Flags().GetString("data")
		metadata, _ := cmd.Flags().GetString("metadata")

		if secretType == "" {
			fmt.Println("Error: must specify the secret type (`--type`)")
			return
		}

		err := api.AddVaultSecret(secretType, data, metadata)
		if err != nil {
			fmt.Println("Error adding secret:", err)
			return
		}

		fmt.Println("The secret was added successfully.!")
	},
}

func init() {
	mainCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("type", "t", "", "Data type (login, text, binary, card)")
	addCmd.Flags().StringP("data", "d", "", "Data to store")
	addCmd.Flags().StringP("metadata", "m", "", "Additional metadata")
}
