package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stakeCmd)
}

var stakeCmd = &cobra.Command{
	Use:   "stake <amount>",
	Short: "Updates the default stake for a node",
	Long:  "Updates the default stake of a node with the specified amount",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stake command")
	},
}
