package main

import "fmt"

var results [][]string

func arrange(array []string, memory []string) [][]string {
	var current string
	if memory == nil {
		memory = []string{}
	}
	for i := 0; i < len(array); i++ {
		current, array = array[i], append(array[:i], array[i+1:]...)
		if len(array) == 0 {
			results = append(results, append(memory, current))
		}
		arrange(append([]string{}, array...), append(memory, current))
		array = append(array[:i], append([]string{current}, array[i:]...)...)
	}
	return results
}

func main() {
	candidates := []string{"Baswedan", "Subianto", "Maharani"}
	chairs := arrange(candidates, nil)
	for _, row := range chairs {
		fmt.Println(row)
	}
}

//In this code, we use a 2D slice to store the results instead of an array. We also use the append function to add elements to a slice and the len function to get the length of a slice.

//We define the arrange function to recursively permute the elements in the array slice and store the results in the results slice. We use the append function to concatenate slices and the [:] notation to create a copy of a slice.

//Finally, we call the arrange function with the candidates slice and print the chairs slice using a for loop.