package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) { //thêm phẩn tử mới vào cuối danh sách
	if lst.tail == nil {
		lst.head = &element[T]{val: v} //nil => v sẽ được gán làm phần tử đầu tiên
		lst.tail = lst.head            //và cuối cùng trong danh sách
	} else {
		lst.tail.next = &element[T]{val: v} //!nil => v sẽ được thêm vào sau phần tử cuôi
		lst.tail = lst.tail.next            //và lst.tail đc cập nhật
	}
}

func (lst *List[T]) All() iter.Seq[T] { //cho phép duyệt qua tất cả các phần tử mà ko cần trả về 1 slice
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func getFib() iter.Seq[int] {
	return func(yield func(int) bool) { //Sử dụng yield để trả về các giá trị fibonacci
		a, b := 1, 1
		for {
			if !yield(a) { //nếu yield trả về false -> dừng lại
				return
			}
			a, b = b, a+b // cập nhật a, b tương ứng
		}
	}
}
func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	for e := range lst.All() {
		fmt.Println(e)
	}
	all := slices.Collect(lst.All()) //thu thập tất cả phần tử từ danh sạch thành 1 slice
	fmt.Println("all:", all)

	for n := range getFib() {
		if n >= 100 {
			break
		}
		fmt.Println(n)
	}
}
