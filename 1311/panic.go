package main

import "os"

func main() {
	panic("a problem") // kích hoạt panic với thông báo. Khi panic được gọi => tất cả các hàm sẽ dừng lại

	_, err := os.Create("/tmp/file")
	if err != nil { //huỷ bỏ nếu trả về 1 lỗi chúng ta k biết cách
		panic(err) // hoặc k muốn xử lý
	}
}
