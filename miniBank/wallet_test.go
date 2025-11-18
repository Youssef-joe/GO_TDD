package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but did not find one")
		}
		// here in got.Error() we convert the error into string
		// in order to compare it with string "want"
		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}
	assertNoError := func (t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s but wanted %s", got , want)
		}
	}

	t.Run("deposit logic", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw logic", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance :=Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}