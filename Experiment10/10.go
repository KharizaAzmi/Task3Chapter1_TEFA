package main

import (
	"fmt"
	"math/rand"
)

func findCompany(array []string, input int) {
	for i := 0; i < len(array); i++ {
		if array[i] == array[input] {
			fmt.Println("Company Found :" + array[input])
			break;
		}
		fmt.Println("Searching Company...", (i+1))
	}
}

func main() {
	telco := []string{"Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon", "Vodafone","Orange", "KDDI", "Telefonica", "T-Mobile"}
	var company int= rand.Intn(len(telco))  
	findCompany(telco[:], company)
	fmt.Println(telco);
}