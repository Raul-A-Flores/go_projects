package main

import "errors"

type Blockchain struct {
	balance int
}

func NewBlockchain() *WellsFargo {
	return &WellsFargo{
		balance: 0,
	}
}

func (b *Blockchain) GetBalance() int {
	return b.balance

}

func (b *Blockchain) Deposit(amount int) {
	b.balance += amount

}

func (b *Blockchain) Withdraw(amount int) error {

	newBalance := b.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}
