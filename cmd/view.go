package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(viewCmd)
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Prints last block in the blockchain",
	Long:  "Prints transaction in the last block of the blockchain along with the validator of it",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("View command")
	},
}
