package main

import "fmt"

type BankAccounts interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	myAccounts := []BankAccounts{
		NewBlockchain(),
		NewWellsFargo(),
	}

	for _, account := range myAccounts {
		account.Deposit(4065)

		if err := account.Withdraw(60); err != nil {
			fmt.Println("Err", err)
		}

		balance := account.GetBalance()
		fmt.Println("Balanace", balance)

	}

	wf := NewWellsFargo()

	wf.Deposit(300)

	currentBalance := wf.GetBalance()

	fmt.Print(currentBalance)

}
