package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(transactionCmd)
}

var transactionCmd = &cobra.Command{
	Use:   "transaction <recipient-id> <coins>|<message>",
	Short: "Create a transaction",
	Long:  "Creates a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Transaction command")
	},
}
