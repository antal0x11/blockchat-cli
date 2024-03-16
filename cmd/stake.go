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
	rootCmd.AddCommand(stakeCmd)
}

var stakeCmd = &cobra.Command{
	Use:   "stake <amount>",
	Short: "Updates the default stake for a node",
	Long:  "Updates the default stake of a node with the specified amount",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Help()
		} else {

			fmt.Println("=> Updating stake.")

			_st := dst.TransactionRequest{}

			stakeAmount, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				cmd.Help()
				return
			}

			_st.TypeOfTransaction = "stake"
			_st.Amount = stakeAmount

			fmt.Printf("=> Setting stake to %.2f.\n", _st.Amount)

			if data, err := json.Marshal(_st); err == nil {

				resp, err := http.Post(os.Getenv("NODE_IP")+"/api/stake", "application/json", bytes.NewBuffer(data))
				if err != nil {
					fmt.Println("=> Network error.")
					return
				}
				defer resp.Body.Close()

				_, err = io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("=> Failed to parse response.")
				}

				if resp.StatusCode == 200 {
					fmt.Println("=> Update stake Status: Sent.")
					fmt.Println("=> The new stake will be set after the next valid block.")
				}

			} else {
				log.Fatal("# Failed to marshal transacion request.")
			}
		}
	},
}
