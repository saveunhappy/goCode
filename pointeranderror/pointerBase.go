package pointeranderror

import "fmt"

// Deposit 定金
func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance 结余
func (w Wallet) Balance() int {
	return w.balance
}

type Wallet struct {
	balance int
}
