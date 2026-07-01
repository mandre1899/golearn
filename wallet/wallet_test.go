package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertError := func (t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("Wanted an error but did't get one.")
		}

		if err.Error() != want {
			t.Errorf("Got error %q want %q", err.Error(), want)
		}
	}

	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "Insufficient funds: Balance(20 BTC), Withdawal(100 BTC)")
		assertBalance(t, wallet, Bitcoin(20))
	})
}
