
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
		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscount(input []string) {
	for _, customer := range input {
		fmt.Println("Sending Promo to " + customer)
	}
	fmt.Println("Finished sending Promo to " + strconv.Itoa(len(input)) + " Customers")

	for _, customer := range input {
		fmt.Println("Sending Discount to " + customer)
	}
	fmt.Println("Finished sending Discount to " + strconv.Itoa(len(input)) + " Customers")
}

func main() {
	data := generateData(1000)
	sendPromoDiscount(data)
}