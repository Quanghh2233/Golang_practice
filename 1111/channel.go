package main

import "fmt"

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() <-chan T { //tạo và trả về 1 channel T
	ch := make(chan T) // để duyệt qua các phần tử trong danh sách
	go func() {        // 1 goroutine đc tạo ra để gửi từng phần tử trong danh sách qua channel ch
		for e := lst.head; e != nil; e = e.next {
			ch <- e.val
		}
		close(ch) // channel ch được đóng, thông báo rằng ko còn phần tử nào nữa
	}()
	return ch
}

func getFib() <-chan int { // tạo ra 1 channel số nguyên
	ch := make(chan int) // và trả về dãy fibonacci thông qua kênh này
	go func() {          //1 goroutine đc sử dụng để tính và gửi các số fibonacci vào channel ch
		a, b := 1, 1
		for {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}
func Collect[T any](ch <-chan T) []T { // nhận 1 channel ch và thu thập tất cả
	var result []T      // các phần tử từ channel này vào 1 slice result
	for v := range ch { // nhận từng phần tử từ kênh ch cho đén khi kênh bị đóng
		result = append(result, v) // các phần tử được thêm vào result và trả về kq cuối cùng
	}
	return result
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := Collect(lst.All())
	fmt.Println("all:", all)

	for n := range getFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

	message := make(chan string) // tạo channel mới, được đặt theo giá trị mà chúng truyền tải

	go func() { message <- "ping" }() //truyền giá trị vài channel bằng cú pháp channel <-

	msg := <-message //cú pháp <-channel nhận 1 giá trị từ channel.
	fmt.Println(msg)
}
