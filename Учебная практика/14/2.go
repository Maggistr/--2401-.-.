package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	b.balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if b.balance < amount {
		return errors.New("недостаточно средств на счете")
	}
	b.balance -= amount
	return nil
}

func (b BankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	account := BankAccount{owner: "Петр Петров", balance: 1000.0}

	err := account.Deposit(500)
	if err != nil {
		fmt.Println("Ошибка пополнения:", err)
	}

	err = account.Withdraw(200)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	err = account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	fmt.Printf("Баланс счёта %s: %.2f\n", account.owner, account.GetBalance())
}
