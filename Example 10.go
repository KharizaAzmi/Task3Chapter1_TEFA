package main

import (
	"fmt"
	"math/rand"
	"time"
)

var telco = []string{
	"Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon",
	"Vodafone", "Orange", "KDDI", "Telefonica", "T-Mobile",
}

func findCompany(array []string, input int) {
	for i := 0; i < len(array); i++ {
		if i == input {
			fmt.Println("Company Found :", array[input])
			break
		}

		fmt.Println("Searching Company...", i+1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	company := rand.Intn(11)
	findCompany(telco, company)
}