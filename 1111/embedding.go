package main

import "fmt"

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// func main() {
// 	co := container{
// 		base: base{
// 			num: 2,
// 		},
// 		str: "some name",
// 	}
// 	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
// 	fmt.Println("also num:", co.base.num)
// 	fmt.Println("describe:", co.describe())

// 	type describer interface {
// 		describe() string
// 	}
// 	var d describer = co
// 	fmt.Println("describer:", d.describe())
// }

type person struct {
	name string
	age  int
}

func (p person) SayHello() {
	fmt.Printf("Hello, my name is %s! I'm %d years old.\n", p.name, p.age)
}

type employee struct {
	person   // nhúng struct person vào employee
	position string
}

func (e employee) Greet() {
	fmt.Println("Welcome!")
}

func main() {
	emp := employee{
		person: person{
			name: "Ron",
			age:  33,
		},
		position: "Leader",
	}
	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.position)
	emp.SayHello()
	emp.Greet()
}
