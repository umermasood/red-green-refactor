package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	if got, want := wallet.Balance(), Bitcoin(10); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
