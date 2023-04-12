package main

import "fmt"

func arrange(array []string) [][]string {
	// if memory == nil {
	// 	memory = []string{}
	// }
	// base case: if only one candidate, return it as a slice of a slice
    if len(array) == 1 {
        return [][]string{{array[0]}}
    }

	
	var results [][]string

	// loop through all candidates and make them the first element of the permutation
    for i, arr := range array {
        // create a new slice with the candidate removed
        current := make([]string, len(array)-1)
        copy(current[:i], array[:i])
        copy(current[i:], array[i+1:])

        // recursive call to generate all permutations of the remaining candidates
        memory := arrange(current)

        // append the current candidate to the beginning of each sub-permutation
        for _, subPerm := range memory {
            result := make([]string, len(subPerm)+1)
            result[0] = arr
            copy(result[1:], subPerm)
            results = append(results, result)
        }
    }

    return results;
}

func main() {
	candidates := []string{"Baswedan", "Subianto", "Maharani", "Ganjar"}
	chairs := arrange(candidates)

	// Print the header row
	fmt.Printf("%-15s %-15s %-15s %-15s\n", "Chair 1", "Chair 2", "Chair 3", "Chair 4")
	fmt.Println("------------------------------------")

	// Print each row of the 2D slice
	for _, row := range chairs {
		fmt.Printf("%-15s %-15s %-15s %-15s\n", row[0], row[1], row[2], row[3])
	}
}
