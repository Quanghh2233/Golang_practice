package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "asdqwe1123423!#@#!@$~:"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) //chuyển chuỗi data thành mảng byte và mã hóa nó bằng chuẩn Base64
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc) //giải mã chuỗi Base64 đã mã hóa sEnc về mảng byte gốc
	fmt.Println(string(sDec))                     //chuyển mảng byte trở lại chuỗi và in kết quả, xác minh rằng chuỗi giải mã đúng với chuỗi ban đầu data
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) //mã hóa chuỗi data sử dụng phiên bản Base64 dành cho URL, an toàn khi truyền qua URL vì nó thay thế ký tự + bằng - và / bằng _
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc) //giải mã
	fmt.Println(string(uDec))
}
