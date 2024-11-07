package main

import "fmt"

func zval(ival int) {
	ival = 0 //=> gán giá trị của ival = 0, tuy nhiên giá trị chỉ ảnh hưởng bản sao của biến trong hàm (1)
}
func zptr(iptr *int) {
	*iptr = 0 //=> sử dụng toán tử * để truy cập giá trị tại địa chỉ hiện tại và đặt giá trị đó về 0 (2)
}

func mod(s []int) { //=>nhận 1 slice kiểu int
	s[0] = 99 //=> thay đổi phần tử đầu tiên của slice thành 99
	s[3] = 10 //=> thay đổi phần tử cuối cùng của slice thành 10
}
func main() {
	i := 6
	fmt.Println("init:", i)

	zval(i)
	fmt.Println("zval:", i) //=> (1) và không ảnh hưởng đến giá trị của biến gốc ở bên ngoài

	zptr(&i)
	fmt.Println("zptr:", i) // (2) => dẫn đến việc giá trị tại địa chỉ của i trong main cũng sẽ bị thay đổi về 0

	fmt.Println("pointer:", &i) //=> xuất địa chỉ của i

	// pointer trong slice
	fmt.Println("Pointer trong slice")
	a := [4]int{0, 22, 44, 88} // tạo slice
	fmt.Println(a)
	mod(a[:]) // =>vì slice này vẫn tham chiếu tới mảng a gốc -> mọi thay đổi sẽ ảnh hưởng đến a
	fmt.Println(a)
}
