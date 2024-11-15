package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.CreateTemp("", "sample") // tạo ra một tệp tạm thời với tên bắt đầu bằng "sample"
	check((err))

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name()) //đảm bảo rằng tệp tạm thời sẽ được xóa khi hàm main() kết thúc

	_, err = f.Write([]byte{1, 2, 3, 4}) //ghi một mảng byte {1, 2, 3, 4} vào tệp tạm thời vừa tạo
	check(err)

	dname, err := os.MkdirTemp("", "sampledir") //tạo một thư mục tạm thời với tên bắt đầu bằng "sampledir"
	check(err)
	fmt.Println("temp dir name:", dname)

	defer os.RemoveAll(dname) //đảm bảo rằng thư mục tạm thời và tất cả nội dung bên trong nó sẽ được xóa khi chương trình kết thúc

	fname := filepath.Join(dname, "file1")        //tạo ra đường dẫn đầy đủ cho tệp "file1" bên trong thư mục tạm thời dname
	err = os.WriteFile(fname, []byte{1, 2}, 0666) //tạo một tệp mới "file1" trong thư mục tạm thời và ghi dữ liệu {1, 2} vào đó (quyền đọc và ghi cho tất cả người dùng).
	check(err)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
