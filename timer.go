package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second)

	<-timer1.C // pause main cho đến khi nhận được dữ liệu từ timer1.C
	fmt.Println("Timer 1 fired")

	time.Sleep(time.Second)
	timer2 := time.NewTimer(time.Millisecond * 100) //timer2 đc tạo với khoảng thời gian 1s. Sau 1s, timer2 sẽ gửi tín hiệu vào time2.C
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	time.Sleep(time.Second) // đủ thời gian để timer2 có thể kích hoạt

	if stop2 := timer2.Stop(); stop2 { // Thử dừng lại timer2 nếu nó chưa kích hoạt
		fmt.Println("Timer 2 stopped")
	} else {
		fmt.Println("Timer 2 had already fired")
	}
	time.Sleep(2 * time.Second)
}
