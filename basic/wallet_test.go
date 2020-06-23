package basic

import "testing"

func assertBalance(t *testing.T, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(BitCoin(10))

		assertBalance(t, wallet, BitCoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, BitCoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})
}
