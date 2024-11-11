package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): //nếu ko có gì đc gửi đến c1 sau 1s, case này đc chọn và việc đọc từ c1 bị bỏ qua
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): //nếu ko có gì đc gửi đến c2 sau 3s, case này đc chọn và việc đọc từ c2 bị bỏ qua
		fmt.Println("timeout 2")
	}
}
