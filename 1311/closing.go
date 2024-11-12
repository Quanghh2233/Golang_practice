package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				// time.Sleep(time.Millisecond * 750)
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all job")
				done <- true // gửi giá trị true vào done để báo hiệu công việc đã hoàn thành
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		// time.Sleep(time.Millisecond * 500)
		fmt.Println("sent job", j)
	}
	close(jobs) //đóng channel jobs
	fmt.Println("sent all jobs")

	<-done //chờ nhận giá trị từ done. Khi true đc gửi vào done, main sẽ tiếp tục thực thi các lệnh sau đó

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
