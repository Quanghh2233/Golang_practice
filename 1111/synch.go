package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) { //nhận 1 kênh done kiểu chan bool => thông báo khi worker hoàn thành công việc
	fmt.Print("Working")
	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Print("done")
	done <- true // gửi giá trị true vào done để báo hiệu rằng công việc đã hoàn thành
}
func main() {
	done := make(chan bool)
	go worker(done)
	<-done //chờ để nhận giá trị từ done. Khi worker gửi true vào done, main sẽ tiếp tục thực thi các câu lệnh sau đó(nếu có)
	time.Sleep(time.Second)
	fmt.Println("\nShutdown?")
}
