package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func multi(a, b, c int) int {
	return a * b * c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = multi(3, 2, 3)
	fmt.Println("3x2x3=", res)

}
