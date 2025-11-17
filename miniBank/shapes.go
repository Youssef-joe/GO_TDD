package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("the address of balance in deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// why did not we make it return (*w).balance ** dereference ** ??
	// the answer is that struct pointers in Go are automatically dereferenced 
	return w.balance
}

// Technically you do not need to 
// change Balance to use a pointer receiver as taking a copy of the balance is fine.
//  However, by convention you should keep
//   your method receiver types the same for consistency.