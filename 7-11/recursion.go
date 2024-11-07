package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(10))
	// tính fibonacci
	var fib func(n int) int
	fib = func(n int) int {
		if n < 5 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(13))
}
