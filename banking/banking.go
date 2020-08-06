package banking

// Account
type Account struct {
	owner   string
	balance int
}

// NewAccount create a new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
