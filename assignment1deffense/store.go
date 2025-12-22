package assignment1deffense

import "fmt"

func CreateStore() *Store {
	return &Store{make(map[uint64]*Order)}
}

type Store struct {
	orders map[uint64]*Order
}

func (s *Store) AddOrder(order *Order) error {
	if _, ok := s.orders[order.ID]; ok {
		return fmt.Errorf("order with ID %d already exists", order.ID)
	}

	s.orders[order.ID] = order
	return nil
}

func (s *Store) AddItem(orderId uint64, item *Item) error {
	if _, ok := s.orders[orderId]; !ok {
		return fmt.Errorf("order with ID %d does not exists", orderId)
	} else if item.Name == "" {
		return fmt.Errorf("item with ID %d has no name", orderId)
	} else if item.Price == 0 {
		return fmt.Errorf("item with ID %d has no price", orderId)
	}

	s.orders[orderId].Items = append(s.orders[orderId].Items, item)
	return nil
}

func (s *Store) RemoveItem(orderId uint64, itemName string) error {
	if _, ok := s.orders[orderId]; !ok {
		return fmt.Errorf("order with ID %d does not exists", orderId)
	}

	var item *Item
	var nArr []*Item
	for index, value := range s.orders[orderId].Items {
		if value.Name != itemName {
			nArr = append(nArr, s.orders[orderId].Items[index])
		} else {
			item = value
		}
	}

	if item == nil {
		return fmt.Errorf("item with ID %d does not exists", orderId)
	}

	s.orders[orderId].Items = nArr
	return nil
}

func (s *Store) ListOrders() []*Order {
	var list []*Order
	for _, order := range s.orders {
		list = append(list, order)
	}
	return list
}
func (s *Store) TotalRevenue() float64 {
	var totalRevenue float64
	for _, order := range s.orders {
		for _, item := range order.Items {
			totalRevenue += item.Price
		}
	}
	return totalRevenue
}
