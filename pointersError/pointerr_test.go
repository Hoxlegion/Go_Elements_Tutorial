package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

    t.Run("deposite", func(t *testing.T) {
        wallet := Wallet{}

        wallet.Deposite(Floenk(10))
        assertBalance(t, wallet, Floenk(10))
    })

    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Floenk(20)}
        err := wallet.Withdraw(Floenk(10))

        assertNoError(t, err)
        assertBalance(t, wallet, Floenk(10))
    })

    t.Run("wathdraw insuffiecient funds", func(t *testing.T){
        startingBalance := Floenk(10)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Floenk(20))

        assertError(t, err, ErrInsuffienctFunds)
        assertBalance(t, wallet, startingBalance)
    })
}

func assertBalance(t testing.TB, w Wallet, want Floenk) {
    t.Helper()
    got := w.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("got an error but didn't want one'")
    }
}

func assertError(t testing.TB, got error, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("wanted an error but didn't get one'")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
