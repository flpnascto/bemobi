package main

import (
	"fmt"

	"github.com/flpnascto/bemobi/internal/entity"
)

func main() {
	account, err := entity.NewAccount(5)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	fmt.Println("Account created successfully")
	fmt.Printf("%+v\n", account)

	firstInvesting, err := entity.NewInvestment(1, 2, 10)
	if err != nil {
		fmt.Println("Error creating firstInvesting:", err)
		return
	}
	fmt.Printf("firstInvesting created successfully %+v\n", firstInvesting)

	secondInvesting, err := entity.NewInvestment(2, 4, 5)
	if err != nil {
		fmt.Println("Error creating secondInvesting:", err)
		return
	}
	fmt.Printf("secondInvesting created successfully %+v\n", secondInvesting)

	thirdInvesting, err := entity.NewInvestment(3, 5, 12)
	if err != nil {
		fmt.Println("Error creating thirdInvesting:", err)
		return
	}
	fmt.Printf("thirdInvesting created successfully %+v\n", thirdInvesting)

	err = account.MakeApplication(*firstInvesting)
	if err != nil {
		fmt.Println("Error adding firstInvesting:", err)
		return
	}

	err = account.MakeApplication(*secondInvesting)
	if err != nil {
		fmt.Println("Error adding secondInvesting:", err)
		return
	}

	err = account.MakeApplication(*thirdInvesting)
	if err != nil {
		fmt.Println("Error adding thirdInvesting:", err)
		return
	}

	fmt.Printf("\nAccount Status\n%+v\n", account)
	fmt.Println("MaxInvestmentValue:", account.MaxInvestmentValue())
}
