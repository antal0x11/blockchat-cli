package main

import (
	"fmt"
	"os"
)

func main() {

	params := len(os.Args)

	switch params {
	case 4:
		switch os.Args[1] {
		case "transaction":
			fmt.Println("Command transaction")
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
