package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/antal0x11/blockchat-cli/dst"
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

		resp, err := http.Get(os.Getenv("NODE_IP") + "/api/balance")
		if err != nil {
			log.Fatal("# bctl: error making http request.")
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("# bctl: error channeling body resposne.")
			}

			var data dst.BalanceResponse
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Fatal("# bctl: error parsing response.")
			}

			fmt.Printf("\n=> Your balance is %.f BCC\n\n", data.Balance)
		} else {
			log.Fatal("# bctl: got an invalid response.")
		}
	},
}
