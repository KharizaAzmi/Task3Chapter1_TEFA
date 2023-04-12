package main

import "fmt"

var counterPlainRecursive int
var counterDynamicProgramming int

func fibonacciPlainRecursive(n int) int {
	counterPlainRecursive++
	if n < 2 {
		return n
	}
	return fibonacciPlainRecursive(n-1) + fibonacciPlainRecursive(n-2)
}

func fibonacciDynamicProgramming(n int) int {
	counterDynamicProgramming++
	if n < 2 {
		return n
	}
	cache := make([]int, n+1)
	cache[0], cache[1] = 0, 1
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}

func main() {
	fibonacciPlainRecursive(20)
	fibonacciDynamicProgramming(20)

	//fmt.Println(fibonacciPlainRecursive(20))
	//fmt.Println(fibonacciDynamicProgramming(20))

	fmt.Println("we did", counterPlainRecursive, "calculation for Plain Recursive")
	fmt.Println("we did", counterDynamicProgramming, "calculation for Dynamic Programming")
}
