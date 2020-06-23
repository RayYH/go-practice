package basic

import (
	"errors"
	"fmt"
)

type BitCoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance BitCoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (bitCoin BitCoin) String() string {
	return fmt.Sprintf("%d BTC", bitCoin)
}

func (wallet *Wallet) Deposit(amount BitCoin) {
	wallet.balance += amount

}

func (wallet *Wallet) Balance() BitCoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount BitCoin) error {

	if amount > wallet.balance {
		return ErrInsufficientFunds
	}

	wallet.balance -= amount

	return nil
}
