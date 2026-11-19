package main

import (
	"fmt"
	"go-functions/internal/payment"
)

func main() {
	amount := int64(10000)
	balance := int64(100002)

	fee, total, err := payment.ProcessWithdrawal(amount, balance)

	if err != nil {
		fmt.Println("withdrawal failed:", err)
		return
	}
	fmt.Printf("Your withdrawal of Kes  %d  was succesful \n Total  fee: %d\n Total amount deducted %d", amount, fee, total)
}
