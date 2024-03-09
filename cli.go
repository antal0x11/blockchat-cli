package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/antal0x11/blockchat-cli/dst"
	"github.com/antal0x11/blockchat-cli/lib"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("# Invalid Configuration.")
	}

	params := len(os.Args)

	switch params {
	case 4:
		switch os.Args[1] {
		case "transaction":

			amount, err := strconv.ParseFloat(os.Args[3], 64)
			if err != nil {
				log.Fatal("# Invalid amount of coins.")
			}

			_t := dst.Transaction{
				SenderAddress:     "Tony",
				RecipientAddress:  os.Args[2],
				TypeOfTransaction: "coins",
				Amount:            amount,
				Message:           "",
				Nonce:             1,
				TransactionId:     "asdf234sdfadf",
				Signature:         "sdfadsf2452342sadf323423",
			}

			lib.TransactionPublisher(_t)
		default:
			help()
		}
	case 3:
		switch os.Args[1] {
		case "stake":
			fmt.Println("Command stake")
		default:
			help()
		}
	case 2:
		switch os.Args[1] {
		case "view":
			fmt.Println("Command view")
		case "balance":
			fmt.Println("Command balance")
		default:
			help()
		}
	default:
		help()
	}
}

func help() {
	fmt.Printf("# CLI for BlockChat\n\n")
	fmt.Println("# ./cli <command> <option> <option>")
	fmt.Printf("\n")
	fmt.Println("# Commands")
	fmt.Printf("------------------------\n\n")
	fmt.Println("# transaction <recipient_address> <coins | message>")
	fmt.Println("# stake <amount>")
	fmt.Println("# view")
	fmt.Println("# balance")
	fmt.Println("# help")
}
