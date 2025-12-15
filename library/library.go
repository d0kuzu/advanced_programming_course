package library

import (
	"errors"
	"fmt"
)

func NewLibrary() *Library {
	return &Library{
		Books: make(map[uint]Book),
	}
}

type Library struct {
	Books map[uint]Book
}

var counter uint = 0

func (l *Library) AddBook(title, author string) {
	l.Books[counter] = Book{ID: counter, Title: title, Author: author, IsBorrowed: false}
	counter++
}

func (l *Library) BorrowBook(id uint) (Book, error) {
	book, ok := l.Books[id]
	if !ok {
		return Book{}, errors.New("book not found")
	}
	book.IsBorrowed = true
	return book, nil
}

func (l *Library) ReturnBook(id uint) error {
	book, ok := l.Books[id]
	if !ok {
		return errors.New("book not found")
	}
	book.IsBorrowed = false
	return nil
}

func (l *Library) ListAvailableBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books available")
	}

	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Println(book.ToString())
		}
	}
}
