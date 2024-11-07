package main

import "fmt"

func vals() (int, int) {
	return 123, 231
}
func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	c, _ := vals()
	fmt.Println(c)
	_, d := vals()
	fmt.Println(d)
}
