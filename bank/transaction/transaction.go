package transaction

import "fmt"

type Transaction struct {
	Type   Types
	Amount float64
}

func (t *Transaction) ToString() string {
	return fmt.Sprintf("Type: %s, Amount: %f", t.Type, t.Amount)
}
