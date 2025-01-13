package main

import "errors"

func (w *Wallet) Deposite(amount Floenk) {
    w.balance += amount
}

func (w *Wallet) Balance() Floenk {
    return w.balance
}

func (w *Wallet) Withdraw(amount Floenk) error {

    if amount > w.balance {
        return errors.New("cannot withdraw, insuffiecient funds")
    }

    w.balance -= amount
    return nil
}
