package main

import (
	"errors"
	"fmt"
)

// Bitcoin is a type
type Bitcoin int

// Wallet is a type
type Wallet struct {
	balance Bitcoin
}

// Stringer is a type
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit deposits into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw bitcoin from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}