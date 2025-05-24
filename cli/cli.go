package cli

import (
	"bufio"
	"fmt"
	"blockchain-project/service"
	"os"
	"strconv"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Mine Block")
		fmt.Println("3. View Blockchain")
		fmt.Println("4. View Wallet Balances")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter sender: ")
			scanner.Scan()
			sender := scanner.Text()

			fmt.Print("Enter receiver: ")
			scanner.Scan()
			receiver := scanner.Text()

			fmt.Print("Enter amount: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(scanner.Text(), 64)

			service.AddTransaction(sender, receiver, amount)

		case "2":
			service.MineBlock()

		case "3":
			for _, b := range service.GetBlockchain() {
				fmt.Printf("Index: %d Hash: %s\n", b.Index, b.Hash)
			}

		case "4":
			for _, w := range service.GetWallets() {
				fmt.Printf("%s: %.2f\n", w.Address, w.Balance)
			}

		case "5":
			fmt.Println("Exiting CLI...")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
