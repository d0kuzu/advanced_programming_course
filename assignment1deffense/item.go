package assignment1deffense

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func (i *Item) ToString() string {
	return fmt.Sprintf("Name: %s Price: %f", i.Name, i.Price)
}
