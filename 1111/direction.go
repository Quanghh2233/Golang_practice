package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, msg string) { //channel chỉ được phép gửi dữ liệu
	pings <- msg //=> lỗi nếu cố tình lấy dữ liệu từ hàm
}
func pong(pings <-chan string, pongs chan<- string) { //channel chỉ được phép nhận dữ liệu => lỗi nếu gửi
	for msg := range pings { // => lỗi nếu cố tình gửi dữ liệu vào hàm
		pongs <- msg
	}
	close(pongs) // đóng pongs sau khi tất cả tin nhắn đã đc chuyển
}
func main() {
	pings := make(chan string, 2)
	pongs := make(chan string, 2)

	ping(pings, "passed message")
	ping(pings, "Creeper")
	close(pings) // đóng pings để báo hiệu k còn tin nhắn mới

	pong(pings, pongs)

	for msg := range pongs {
		time.Sleep(time.Second)
		fmt.Println(msg)
	}

}
