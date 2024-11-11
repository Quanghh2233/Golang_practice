package main

import (
	"fmt"
	"time"
)

func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // Mỗi loop, select sẽ lựa chọn 1 trong các case có sẵn
		case c <- x: //gửi giá trị của x vào c, sau đó cập nhật x, y theo quy tắc fibonacci
			x, y = y, x+y
		case <-quit: //Khi có giá trị nhận từ quit, in ra "quit" và thoát khỏi hàm bằng return
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 400)
			fmt.Println(<-c)
		}
		quit <- 0 //sau khi loop đến điều kiện, gửi 1 giá trị vào quit để kích hoạt case <-quit, khiến nó thoát loop và kết thúc
	}()
	fibo(c, quit)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
