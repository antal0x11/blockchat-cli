package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/antal0x11/blockchat-cli/dst"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(transactionCmd)
}

var transactionCmd = &cobra.Command{
	Use:   `transaction <recipient-id> <coins>|"<message>"`,
	Short: "Create a transaction",
	Long:  "Creates a transaction",
	Run: func(cmd *cobra.Command, args []string) {

		_req := dst.TransactionRequest{}

		if len(args) != 2 {
			cmd.Help()
		} else {

			fmt.Println("=> Creating transaction.")

			recipientID, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				cmd.Help()
				return
			}

			_req.RecipientAddressID = uint32(recipientID)

			context, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				_req.TypeOfTransaction = "message"
				_req.Message = args[1]
			} else {
				_req.TypeOfTransaction = "coins"
				_req.Amount = context
			}

			switch _req.TypeOfTransaction {
			case "coins":
				fmt.Printf("=> Transaction of type: %s with context: %.2f\n", _req.TypeOfTransaction, _req.Amount)
			case "message":
				fmt.Printf(`=> Transaction of type: %s with context: "%s"`, _req.TypeOfTransaction, _req.Message)
				fmt.Println()
			}

			if data, err := json.Marshal(_req); err == nil {

				resp, err := http.Post(os.Getenv("NODE_IP")+"/transaction", "application/json", bytes.NewBuffer(data))
				if err != nil {
					fmt.Println("=> Network error")
					return
				}
				defer resp.Body.Close()

				_, err = io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("=> Failed to parse response")
				}

				if resp.StatusCode == 200 {
					fmt.Println("=> Transaction Status: Sent.")
				}

			} else {
				log.Fatal("# Failed to marshal transacion request.")
			}

		}
	},
}
