package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Account struct {
	ID      int
	Owner   string
	Balance float64
}

type Bank struct {
	Name     string
	Accounts []Account
}

var (
	ErrInvalidAmount   = errors.New("invalid amount. Must be > 0")
	ErrAccountNotFound = errors.New("account not found in bank")
	ErrNotEnoughFunds  = errors.New(
		"not enough money on the balance sheet to complete this operation",
	)
	ErrAccountIsAlreadyExists = errors.New("Account is already exists")
)

func (b Bank) getAccountByID(id int) (*Account, error) {
	for i := range b.Accounts {
		if b.Accounts[i].ID == id {
			return &b.Accounts[i], nil
		}
	}

	return nil, fmt.Errorf("account %d: %w", id, ErrAccountNotFound)
}

func (b *Bank) addAccount(acc Account) (*Account, error) {
	_, err := b.getAccountByID(acc.ID)
	if err == nil {
		return nil, ErrAccountIsAlreadyExists
	}

	b.Accounts = append(b.Accounts, acc)

	return &b.Accounts[len(b.Accounts)-1], nil
}

func (b *Bank) topUpAccount(id int, amount float64) (float64, error) {
	if amount <= 0 {
		return 0, ErrInvalidAmount
	}

	account, err := b.getAccountByID(id)
	if err != nil {
		return 0, ErrAccountNotFound
	}

	account.Balance += amount

	return account.Balance, nil
}

func (b *Bank) withdrawFromAccount(id int, amount float64) (float64, error) {
	if amount <= 0 {
		return 0, ErrInvalidAmount
	}

	account, err := b.getAccountByID(id)
	if err != nil {
		return 0, ErrAccountNotFound
	}

	if account.Balance < amount {
		return 0, ErrNotEnoughFunds
	}

	account.Balance -= amount

	return account.Balance, nil
}

func (b *Bank) transferBetweenAccounts(
	idFrom int,
	idTo int,
	amount float64,
) error {
	_, err := b.getAccountByID(idFrom)
	if err != nil {
		return ErrAccountNotFound
	}

	_, err = b.getAccountByID(idTo)
	if err != nil {
		return ErrAccountNotFound
	}

	if _, err := b.withdrawFromAccount(idFrom, amount); err != nil {
		return err
	}
	if _, err := b.topUpAccount(idTo, amount); err != nil {
		return err
	}

	return nil
}

func (b Bank) printAllAccounts() {
	for i := range b.Accounts {
		acc := b.Accounts[i]

		fmt.Println(
			"ID: " + strconv.Itoa(acc.ID) + "\n" +
				"Owner: " + acc.Owner + "\n" +
				"Balance: " + strconv.FormatFloat(acc.Balance, 'f', 2, 64),
		)

	}
}

func main() {
	sberbank := Bank{
		Name: "sberbank",
	}

	account1 := Account{
		ID:      1,
		Owner:   "Elon Musk",
		Balance: 100_000_000,
	}

	account2 := Account{
		ID:      2,
		Owner:   "Poor Student",
		Balance: 5,
	}

	accounts := []Account{account1, account2}

	sberbank.Accounts = append(sberbank.Accounts, accounts...)

	sberbank.printAllAccounts()

	sberbank.transferBetweenAccounts(account1.ID, account2.ID, 1000)

	sberbank.printAllAccounts()

	_, err := sberbank.getAccountByID(3)
	fmt.Println(errors.Is(err, ErrAccountNotFound))
}
