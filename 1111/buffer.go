package main

import (
	"fmt"
	"time"
)

func DemoChannelString() {
	pipe := make(chan string, 2)
	go func() {
		pipe <- "water"
		pipe <- "oil"
		close(pipe)
	}()
	for receiver := range pipe {
		fmt.Println(receiver)
		time.Sleep(time.Millisecond * 500)
	}
}

func write(ch chan int) { //nhận 1 channel ch kiểu chan int
	for i := 0; i < 5; i++ {
		ch <- i //mỗi lần loop gửi giá trị i vào ch
		fmt.Println("Successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	ch := make(chan int, 2) //tạo channel có buffer = 2 => có thể lưu trữ tối đa 2 value mà k cần goroutine khác đọc ngay lập tức
	go write(ch)            // tạo goroutine mới và gọi hàm write, truyền ch vào goroutine này
	time.Sleep(time.Millisecond * 500)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(400 * time.Millisecond)
	}

	DemoChannelString()

	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	time.Sleep(time.Millisecond * 500)
	fmt.Println(<-messages)
}
