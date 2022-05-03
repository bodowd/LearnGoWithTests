package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// because we are mutating state with this function, need pointer to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// var keyword allows us to define values global to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
