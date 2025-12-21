package bank

import (
	transaction2 "advance/assignment1/bank/transaction"
	"errors"
)

func GetAccount() *Account {
	return &Account{balance: 0}
}

type Account struct {
	balance      float64
	transactions []*transaction2.Transaction
}

func (account *Account) Deposit(amount float64) {
	account.balance += amount
	account.transactions = append(account.transactions, &transaction2.Transaction{Type: transaction2.Deposit, Amount: amount})
}

func (account *Account) Withdraw(amount float64) error {
	if account.balance < amount {
		return errors.New("not enough money")
	}

	account.balance -= amount
	account.transactions = append(account.transactions, &transaction2.Transaction{Type: transaction2.Withdraw, Amount: amount})
	return nil
}

func (account *Account) GetBalance() float64 {
	return account.balance
}
