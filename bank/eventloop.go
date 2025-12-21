package bank

import (
	"advance/helpers/console"
	"advance/helpers/reader"
	"fmt"
)

func Start() {
	fmt.Println("Bank commands")
	fmt.Println("/d - deposit")
	fmt.Println("/w - withdraw")
	fmt.Println("/b - get balance")
	fmt.Println("/l - list transactions")
	fmt.Println("/q - quit")
	fmt.Println("_________________")

	account := GetAccount()
	r := reader.NewReader()

	for {
		var action string
		r.ReadValue(&action)

		switch action {
		case "/q":
			fmt.Println("Quiting...")
			return

		case "/d":
			fmt.Println("Write amount")
			var amount float64
			r.ReadValue(&amount)

			account.Deposit(amount)
			fmt.Println("Account deposited")

		case "/w":
			fmt.Println("Write amount")
			var amount float64
			r.ReadValue(&amount)

			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Account withdrawn")

		case "/b":
			fmt.Printf("Balance: %f\n", account.GetBalance())

		case "/l":
			for _, value := range account.transactions {
				fmt.Println(value.ToString())
			}
		}

		console.ClearScreen()
	}
}
