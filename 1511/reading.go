package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("/tmp/dat") //đọc toàn bộ nội dung của tệp /tmp/dat và lưu vào dat.
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat") //mở tệp ở chế độ chỉ đọc và trả về con trỏ tệp f
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1) //đọc tối đa 5 byte vào mảng b1.
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart) //di chuyển con trỏ tệp đến vị trí thứ 6 từ đầu.
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(4, io.SeekCurrent) //di chuyển con trỏ tệp thêm 4 byte từ vị trí hiện tại.
	check(err)

	_, err = f.Seek(10, io.SeekEnd) //di chuyển con trỏ ngược lại 10 byte từ cuối tệp (nma để -10 lại lỗi??? đhs)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2) //đảm bảo đọc ít nhất 2 byte vào b3
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart) //đặt lại con trỏ tệp về đầu tệp.
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) //đọc trước 5 byte mà không di chuyển con trỏ đọc.
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
