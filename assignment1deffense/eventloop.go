package assignment1deffense

import (
	"advance/assignment1/helpers/reader"
	"fmt"
)

func Start() {
	fmt.Println("Commands:")
	fmt.Println("1 - create order")
	fmt.Println("2 - add item to order")
	fmt.Println("3 - remove item from order")
	fmt.Println("4 - list orders")
	fmt.Println("5 - show total revenue")
	fmt.Println("6 - exit")

	r := reader.NewReader()
	store := CreateStore()

	for {
		var action string
		r.ReadValue(&action)

		switch action {
		case "6":
			return

		case "1":
			fmt.Println("Enter order ID")
			var id uint64
			r.ReadValue(&id)

			err := store.AddOrder(&Order{ID: id})
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Order created")

		case "2":
			fmt.Println("Enter order ID")
			var id uint64
			r.ReadValue(&id)

			fmt.Println("Enter item name")
			var name string
			_ = r.ReadString(&name, '\n')

			fmt.Println("Enter item price")
			var price float64
			r.ReadValue(&price)

			err := store.AddItem(id, &Item{name, price})
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Item added")

		case "3":
			fmt.Println("Enter order ID")
			var id uint64
			r.ReadValue(&id)

			fmt.Println("Enter item name")
			var name string
			_ = r.ReadString(&name, '\n')

			err := store.RemoveItem(id, name)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Item removed")

		case "4":

			for _, item := range store.ListOrders() {
				fmt.Println(item.ToString())
			}

		case "5":
			amount := store.TotalRevenue()

			fmt.Println("Total Revenue amount:", amount)
		}
	}
}
