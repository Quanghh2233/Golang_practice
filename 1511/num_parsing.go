package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64) //chuyển chuỗi "1.234" thành số thực float64 với độ chính xác 64 bit.
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) //chuyển chuỗi "123" thành số nguyên int64. Tham số 0 cho phép tự động xác định cơ số (nếu chuỗi có "0x" thì hiểu là cơ số 16)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) // chuyển chuỗi "0x1c8" (số hệ 16) thành số nguyên int64.
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) // chuyển chuỗi "789" thành số nguyên không dấu uint64.
	fmt.Println(u)

	k, _ := strconv.Atoi("135") // chuyển chuỗi "135" thành int (mặc định là int trong hệ thống).
	fmt.Println(k)

	_, e := strconv.Atoi("wat") //hông thể chuyển đổi chuỗi "wat" thành số nguyên, vì vậy nó trả về lỗi e.
	fmt.Println(e)
}
