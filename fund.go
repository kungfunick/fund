package funding

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worring about
	// wheteher it's on the stack or heap
	return &Fund{
		balance: initialBalance,
	}
}

// Method starts with a reciever - Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
