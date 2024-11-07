package main

import "fmt"

type person struct {
	name string
	age  int
}

func newP(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p //con trỏ đến person
}
func main() {
	fmt.Println(person{"Brian", 20}) //tạo 1 struct mới

	fmt.Println(person{name: "Alice", age: 30}) // có thể đặt tên cho các trường khi tạo struct mới

	fmt.Println(person{name: "Bruce"}) //trường bị bỏ qua sẽ được gán giá trị = 0

	fmt.Println(&person{name: "Annie", age: 45}) //& tạo 1 con trỏ đến thẳng person

	fmt.Println(newP("Josh")) // tạo ra 1 person mới, sau đó trả về con trỏ đến đối tượng này

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // truy cập vào struct bằng dấu chấm

	sp := &s
	fmt.Println(sp.age) //sử dụng dot với struct pointer -> con trỏ sẽ tự động huỷ tham chiếu đến đối tượng và lấy giá trị của trường age

	sp.age = 51
	fmt.Println(sp.age) //struct có thể thay đổi

	fmt.Println(s)

	dog := struct { // Struct đùng cho 1 giá trị duy nhất => k cần thiết phải đặt tên
		name   string // Giá trị có thể ở kiểu struct ẩn danh
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
