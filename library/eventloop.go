package library

import (
	"advance/helpers/console"
	"advance/helpers/reader"
	"fmt"
)

func Start() {
	fmt.Println("Library commands")
	fmt.Println("/a - add a book")
	fmt.Println("/b - borrow a book")
	fmt.Println("/r - return a book")
	fmt.Println("/l - list all books")
	fmt.Println("/q - quit")
	fmt.Println("_________________")

	lib := NewLibrary()
	r := reader.NewReader()

	for {
		var action string
		r.ReadValue(&action)

		switch action {
		case "/q":
			fmt.Println("Quiting...")
			return

		case "/a":
			var title, author string

			fmt.Println("Enter book title")
			_ = r.ReadString(&title, '\n')

			fmt.Println("Enter book author")
			_ = r.ReadString(&author, '\n')

			lib.AddBook(title, author)
			fmt.Println("Book added successfully")

		case "/b":
			fmt.Println("Enter book ID")

			var id uint
			r.ReadValue(&id)

			book, err := lib.BorrowBook(id)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println("Borrowed book: " + book.Title)

		case "/r":
			fmt.Println("Enter book ID")

			var id uint
			r.ReadValue(&id)

			err := lib.ReturnBook(id)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println("Book returned successfully")

		case "/l":
			lib.ListAvailableBooks()

		default:
			fmt.Println("Unknown action: " + action)
		}

		console.ClearScreen()
	}
}
