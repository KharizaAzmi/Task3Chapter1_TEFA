package main

import (
	"log"
	"time"
)

const company = "telkom"

func findCompany(array []string){
	arrLength := len(array)
	start := time.Now()

    for i := 0; i < arrLength; i++ {
		if(array[i] == "telkom"){
			log.Printf("Company Found!");
		}
	}

    elapsed := time.Since(start)
    log.Printf("The performance is %s", elapsed)
}

func main(){
	largeCompanyName := make([]string, 1000)
    for i := range largeCompanyName {
        largeCompanyName[i] = company
    }
    findCompany(largeCompanyName)
}