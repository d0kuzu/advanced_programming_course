package library

import "fmt"

type Book struct {
	ID         uint
	Title      string
	Author     string
	IsBorrowed bool
}

func (b Book) ToString() string {
	return fmt.Sprintf("ID: %d, Title: %s, Author: %s, is borrowed: %t", b.ID, b.Title, b.Author, b.IsBorrowed)
}
