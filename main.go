package main

import (
	"fmt"

	"github.com/sy9777m/learnGo/banking"
)

func main() {
	account := banking.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}

}
