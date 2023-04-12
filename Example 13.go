package main

import (
	"fmt"
	"strconv"
)

func generateData(n int) []string {
	baseNumber := "082"
	var customers []string
	var mobileNumber string

	for i := 0; i < n; i++ {
		mobileNumber = baseNumber + strconv.Itoa(i+1)
		for len(mobileNumber) < 12 {
			mobileNumber = "0" + mobileNumber
		}
		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscount(input1 []string, input2 []string) {
	for i := 0; i < len(input1); i++ {
		fmt.Println("Sending Promo to " + input1[i])
	}
	fmt.Println("Its Finished to send Promo to " + strconv.Itoa(len(input1)) + " Customers")

	for i := 0; i < len(input2); i++ {
		fmt.Println("Sending Discount to " + input2[i])
	}
	fmt.Println("Its Finished to send Discount to " + strconv.Itoa(len(input2)) + " Customers")
}

func main() {
	dataPromo := generateData(100000)
	dataDiscount := generateData(1000)

	sendPromoDiscount(dataPromo, dataDiscount)
}