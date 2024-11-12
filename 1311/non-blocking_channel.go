package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string, 1)
	signal := make(chan string, 1)

	go func() {
		time.Sleep(500 * time.Millisecond)
		signal <- "start"
	}()

	go func() {
		time.Sleep(1000 * time.Millisecond)
		message <- "hi"
	}()

	msg := <-message
	sig := <-signal

	// Đợi goroutines gửi dữ liệu và nhận dữ liệu từ các kênh
	time.Sleep(1 * time.Second)
	select {
	case msg := <-message:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("received message", msg)
	case sig := <-signal:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("received signal", sig)
	default:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("no message received")
	}

	// Gửi lại dữ liệu từ các kênh nếu có dữ liệu
	select {
	case message <- msg:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("sent message:", msg)
	case signal <- sig:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("sent signal:", sig)
	default:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("no message sent")
	}

	// Nhận từ các kênh lần cuối nếu có dữ liệu
	select {
	case msg := <-message:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("received message:", msg)
	case sig := <-signal:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("received signal:", sig)
	default:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("no activity")
	}
}
