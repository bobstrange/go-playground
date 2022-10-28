package main

import (
	"encoding/json"
	"fmt"
)

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c CreditCard) String() string {
	// mask output
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	// mask output
	return "xxxx-xxxx-xxxx-xxxx"
}

func main() {
	c := ConfidentialCustomer{
		CustomerID: 100,
		CreditCard: "1111-1111-1111-1111",
	}

	fmt.Println(c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	// GoStringer https://budougumi0617.github.io/2019/10/12/confirm-print-with-fmt-interfaces/
	fmt.Printf("%#v\n", c)

	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes))

	// 結果
	// {100 xxxx-xxxx-xxxx-xxxx}
	// {100 xxxx-xxxx-xxxx-xxxx}
	// {CustomerID:100 CreditCard:xxxx-xxxx-xxxx-xxxx}
	// main.ConfidentialCustomer{CustomerID:100, CreditCard:xxxx-xxxx-xxxx-xxxx}
	// JSON:  {"CustomerID":100,"CreditCard":"1111-1111-1111-1111"}

}
