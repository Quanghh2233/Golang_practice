package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	h := sha256.New()  //tạo hash mới
	h.Write([]byte(s)) //chuyển chuỗi s thành một mảng byte ([]byte) và ghi vào đối tượng SHA-256.
	bs := h.Sum(nil)   //trả về giá trị băm của dữ liệu đã được ghi vào. Kết quả là một mảng byte đại diện cho băm SHA-256

	fmt.Println(s)
	fmt.Printf("%x\n", bs) //In kết quả băm dạng hexa
}
