package wallet

import (
	"fmt"
	"errors"
)

// i have named it first insufficientFundsError but the error was 
// error var insufficientFundsError should have name of the form errFoo 
// so that's way i've changed it to ErrInsufficientFunds to have that Err first
var ErrInsufficientFunds  = errors.New("cannot withdraw, insufficient funds")


//fuck primitive types ima making ma own one here
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

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("the address of balance in deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// why did not we make it return (*w).balance ** dereference ** ??
	// the answer is that struct pointers in Go are automatically dereferenced 
	return w.balance
}

// Technically you do not need to 
// change Balance to use a pointer receiver as taking a copy of the balance is fine.
//  However, by convention you should keep
//   your method receiver types the same for consistency.





func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

