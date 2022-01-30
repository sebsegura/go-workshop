package main

import (
	"errors"
	"fmt"
)

// Account
type Account interface {
	CheckAccount(name string) error
}

type account struct {
	name string
}

func newAccount(name string) Account {
	return &account{
		name,
	}
}

func (a *account) CheckAccount(name string) error {
	if a.name != name {
		return errors.New("account name is incorrect")
	}
	fmt.Println("Account verified!")
	return nil
}

// Transaction
type Transaction interface {
	Send()
}

type transaction struct {
	owner  Account
	amount int
}

func newTrx(owner Account, amount int) Transaction {
	return &transaction{
		owner,
		amount,
	}
}

func (trx *transaction) Send() {
	fmt.Println("Transaction sent!")
}

// Facade
type Facade interface {
	Process(id string) error
}

type facade struct {
	account Account
	trx     Transaction
}

func newFacade(accountID string, amount int) Facade {
	account := newAccount(accountID)
	trx := newTrx(account, 3)

	return &facade{
		account,
		trx,
	}
}

func (f *facade) Process(id string) error {
	fmt.Println("Verifying account...")
	err := f.account.CheckAccount(id)
	if err != nil {
		return err
	}

	f.trx.Send()

	return nil
}

// Main
func main() {
	f := newFacade("test", 3)
	err := f.Process("test")
	if err != nil {
		fmt.Println(err)
	}
}
