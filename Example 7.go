package main

import "fmt"

func main() {
    foods := []string{"Sate", "Bakso", "Dimsum", "Rames"}
    drinks := []string{"Jeruk", "Teh", "Kelapa", "Cendol"}
    counter := 0

    for i := 0; i < len(foods); i++ {
        for j := 0; j < len(drinks); j++ {
            counter++
            fmt.Printf("Menu %d: %s dan %s\n", counter, foods[i], drinks[j])
        }
    }
}