package main

import (
	"fmt"
	"time"
)

const company = "telkom"

func findCompany(array []string) {
    tx := time.Now()
    for i := 0; i < len(array); i++ {
        if array[i] == "telkom" {
            fmt.Println("Company Found!")
        }
    }
    ty := time.Now()
    fmt.Println("The performance is", ty.Sub(tx).Milliseconds(), "ms")
}

func main() {
    largeCompanyName := make([]string, 1000)
    for i := range largeCompanyName {
        largeCompanyName[i] = company
    }
    findCompany(largeCompanyName)
}

//In this code, we use the time package to measure the execution time instead of the performance object in JavaScript. We also use a slice instead of an array since slices are more commonly used in Go. Additionally, we initialize the largeCompanyName slice using the make function and loop through the slice using the range keyword in Go.