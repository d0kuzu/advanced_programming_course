package assignment1deffense

import "fmt"

type Order struct {
	ID    uint64
	Items []*Item
}

func (o *Order) ToString() string {
	itemsStr := ""
	for i, item := range o.Items {
		if i > 0 {
			itemsStr += ", "
		}
		itemsStr += fmt.Sprint(item)
	}
	return fmt.Sprintf("ID: %d Items: [%s]", o.ID, itemsStr)
}
