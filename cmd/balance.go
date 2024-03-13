package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceCmd)
}

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Prints balance",
	Long:  "Prints the balance of the node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Balance command")
	},
}
