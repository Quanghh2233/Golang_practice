package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "orange", "kiwi", "berry"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type person struct {
		name string
		age  int
	}

	people := []person{
		person{name: "Jax", age: 37},
		person{name: "CJ", age: 42},
		person{name: "Louis", age: 28},
		person{name: "Dave", age: 33},
	}
	slices.SortFunc(people,
		func(a, b person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
