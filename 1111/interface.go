package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type student struct {
	Human         // kế thừa từ Human
	school string // bổ sung các thuộc tính riêng
	loan   float32
}
type Employee struct {
	Human          // kế thừa từ Human
	company string // bổ sung các thuộc tính riêng
	money   float32
}
type action interface {
	SayHi() //interface gồm 2 method SayHi() và Hobby()
	Hobby(hobby string)
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s. You can call me on %s\n", h.name, h.phone)
}

func (h Human) Hobby(hobby string) {
	fmt.Println("I like playing", hobby)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s. I work at %s. \nCall me on %s\n", e.name, e.company, e.phone)
}

// Sử dụng con trỏ
type Des interface {
	Des()
}

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Des() { //Sử dụng giá trị receiver
	fmt.Printf("%s is %d years old, living in %s\n", p.name, p.age, p.address)
}

type Addr struct {
	state   string
	country string
}

func (p *Person) Upd(newage int, newAddr string) {
	p.age = newage
	p.address = newAddr
}

func (a *Addr) Des() { //Sử dụng con trỏ receiver
	fmt.Printf("Capital %s Country %s\n", a.state, a.country)
}

func main() {
	Antony := student{Human{"Antony", 25, "222-222-XXX"}, "MIT", 0.00}
	Marcus := student{Human{"Marcus", 26, "111-222-XXX"}, "Harvard", 100}
	John := Employee{Human{"John", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Garry := Employee{Human{"Garry", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	var i action
	i = Antony
	fmt.Println("This is Antony, a student:")
	i.SayHi() // triển khai interface action bằng cách định nghĩa method SayHi() và Hobby()
	i.Hobby("chess")

	i = Marcus
	fmt.Println("This is Marcus, a student:")
	i.SayHi() // triển khai interface action bằng cách định nghĩa method SayHi() và Hobby()
	i.Hobby("skateboard")

	i = Garry
	fmt.Println("This is Garry, an employee:")
	i.SayHi()

	fmt.Println("\nUsing slice")
	x := make([]action, 3)
	x[0], x[1], x[2] = Garry, John, Marcus
	for _, value := range x {
		value.SayHi()
	}

	fmt.Println("\nUsing pointer")
	var d1 Des
	p1 := Person{"Zara", 25, "Sydney"}
	d1 = p1
	d1.Des()
	p2 := Person{"Joshep", 45, "Hawaii"}
	d1 = &p2
	d1.Des()

	p2.Upd(50, "Bangkok")
	d1.Des()

	var d2 Des
	a := Addr{"Moscow", "Russia"}
	d2 = &a
	d2.Des()
}
