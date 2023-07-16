package pointeranderror

// Deposit 定金
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance 结余

func (w *Wallet) Balance() int {
	return w.balance
}

type Wallet struct {
	balance int
}
