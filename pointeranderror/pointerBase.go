package pointeranderror

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return 0
}

type Wallet struct {
	balance int
}
