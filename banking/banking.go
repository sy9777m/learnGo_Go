package banking

import "errors"

// Account create a account
type Account struct {
	owner   string
	balance int
}

// NewAccount create a new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw amount of your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw. You are poor.")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner Owner of the account
func (a *Account) ChangeOwner(owner string) {
	a.owner = owner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return "whatever you want"
}
