package cmd

import (
	"encoding/base64"
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
	rootCmd.AddCommand(viewCmd)
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Prints transactions in the last block of the blockchain",
	Long:  "Prints transaction in the last block of the blockchain along with the validator of it",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := http.Get(os.Getenv("NODE_IP") + "/api/view")
		if err != nil {
			log.Fatal("# bctl: error making http request.")
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {

			body, err := io.ReadAll(resp.Body)

			if err != nil {
				log.Fatal("# bctl: error channeling body resposne.")
			}

			var data dst.Block
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Fatal("# bctl: error parsing response.")
			}

			fmt.Printf("\n=> Validator: %s\n", data.Validator)
			fmt.Printf("-------------------------------------------\n")
			for _, _t := range data.Transactions {
				fmt.Printf("=> Sender Address: %s\n", _t.SenderAddress)
				fmt.Printf("=> Recipient Address: %s", _t.RecipientAddress)
				fmt.Printf("=> Type of Transaction: %s\n", _t.TypeOfTransaction)
				fmt.Printf("=> Amount: %.2f\n", _t.Amount)
				fmt.Printf("=> Message: %s\n", _t.Message)
				fmt.Printf("=> Fee: %.2f\n", _t.Fee)
				fmt.Printf("=> Nonce: %d\n", _t.Nonce)
				fmt.Printf("=> Transaction ID: %s\n", _t.TransactionId)
				fmt.Printf("=> Signature: %s\n\n", base64.StdEncoding.EncodeToString(_t.Signature))
			}
		} else {
			log.Fatal("# bctl: got an invalid response.")
		}
	},
}
