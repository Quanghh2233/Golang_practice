package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644) //ghi dữ liệu d1 vào tệp /tmp/dat1 với quyền truy cập là 0644 (chỉ định tệp tin có quyền đọc-ghi cho người sở hữu và chỉ đọc cho các nhóm khác)
	check(err)

	f, err := os.Create("/tmp/dat2") // tạo tệp /tmp/dat2 và mở tệp này ở chế độ ghi
	check(err)

	defer f.Close() //đảm bảo rằng tệp sẽ được đóng sau khi main kết thúc

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() //đảm bảo rằng mọi dữ liệu được ghi vào tệp từ bộ đệm hệ điều hành sẽ được lưu thực tế vào đĩa

	w := bufio.NewWriter(f) //tạo một bộ ghi đệm w để ghi vào tệp f
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() //ghi mọi dữ liệu còn trong bộ đệm w vào tệp f
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
