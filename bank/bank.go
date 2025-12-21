package bank

import (
	"advance/bank/transaction"
	"errors"
)

func GetAccount() *Account {
	return &Account{balance: 0}
}

type Account struct {
	balance      float64
	transactions []*transaction.Transaction
}

func (account *Account) Deposit(amount float64) {
	account.balance += amount
	account.transactions = append(account.transactions, &transaction.Transaction{Type: transaction.Deposit, Amount: amount})
}

func (account *Account) Withdraw(amount float64) error {
	if account.balance < amount {
		return errors.New("not enough money")
	}

	account.balance -= amount
	account.transactions = append(account.transactions, &transaction.Transaction{Type: transaction.Withdraw, Amount: amount})
	return nil
}

func (account *Account) GetBalance() float64 {
	return account.balance
}
