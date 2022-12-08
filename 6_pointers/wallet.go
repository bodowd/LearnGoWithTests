package main

import (
	"errors"
	"fmt"
)

// !!!! in Go, when you call a function or a method, the arguments are copied
// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
// to keep state, we need to find the address of the bit of memory where wallet is stored

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// take a pointer to the wallet instead of a copy, so that we can change the original values within it
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	// struct pointers are automatically dereferenced in Go so we don't need to
	// write (*w).balance although that is valid as well
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
