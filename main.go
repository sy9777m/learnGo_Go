package main

import (
	"fmt"

	"github.com/sy9777m/learnGo/banking"
)

func main() {
	account := banking.NewAccount("nico")
	fmt.Println(account)
}
