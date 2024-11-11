package main

import (
	"fmt"
	"time"
)

func number() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

}
func alphabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func f(from string) {
	for i := 0; i <= 3; i++ {
		time.Sleep(325 * time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func hello() {
	time.Sleep(475 * time.Millisecond)
	fmt.Println("Hello world")
}

func main() {
	go number()   //goroutine gọi number
	go alphabet() // goroutine gọi alphabet
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nTerminated")

	f("direct")       //gọi hàm trực tiếp
	go f("goroutine") // goroutine gọi f(goroutine)

	go func(msg string) {
		fmt.Println(msg)

	}("going") //kết quả từ goroutine chạy hàm ẩn danh

	time.Sleep(1500 * time.Millisecond)
	fmt.Println("done")

	go hello() //goroutine gọi hello
	time.Sleep(time.Second)
	fmt.Println("main function")
}
