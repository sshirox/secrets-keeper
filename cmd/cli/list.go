package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sshirox/secrets-keeper/internal/cliclient/api"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets",
	Run: func(cmd *cobra.Command, args []string) {
		secrets, err := api.GetVaultSecrets()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, secret := range secrets {
			fmt.Printf("ID: %s\nType: %s\nData: %s\nMetadata: %s\n\n",
				secret["id"], secret["type"], secret["data"], secret["metadata"])
		}
	},
}

func init() {
	mainCmd.AddCommand(listCmd)
}
