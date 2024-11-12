package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range request { //duyệt từng yêu cầu từ request cho đến khi kênh bị đóng
		<-limiter //đảm bảo chỉ 1 yêu cầu được xử lý mỗi 200ms
		fmt.Println("request", req, time.Now())
	}
	time.Sleep(time.Second)

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t // thêm vào 1 giá trị mới mỗi 200ms
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest { // in ra giá trị cùng thời gian hiện tại, sử dụng burstyLimiter để giới hạn tốc độ xử lý
		<-burstyLimiter // cho phép tối đa 3 req liên tiếp mà k cần chờ, sau đó giới hạn mỗi 200ms
		fmt.Println("request", req, time.Now())
	}
}
