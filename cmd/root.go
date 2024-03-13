package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bctl",
	Short: "bctl is a command line interface to interact with blockchat",
	Long:  "\n- bctl is a command line interface to interact with blockchat,\n- a simple platform to exchange coins and messages, based on \n- blockchain technology",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableFlagsInUseLine = true
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
