package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "3"

	close(queue)

	for elem := range queue { // dùng loop để nhận giá trị từ queue
		fmt.Println(elem)
	}

}
