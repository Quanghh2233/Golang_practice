package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename") //dùng để nối các thành phần của đường dẫn vào với nhau, tự động thêm dấu gạch chéo (/ hoặc \ tùy hệ điều hành) giữa các phần
	fmt.Println("p:", p)
	fmt.Println(filepath.Join("dir1//", "filename"))       //Nếu sử dung // => bỏ qua và chỉ sử dụng một dấu gạch chéo
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) //có thể được giải thích là "quay lại thư mục cha của dir1 và sau đó vào lại dir1"
	fmt.Println("Dir(p):", filepath.Dir(p))                //sẽ trả về thư mục chứa tệp trong đường dẫn
	fmt.Println("Base(p):", filepath.Base(p))              //sẽ trả về phần cuối cùng của đường dẫn (tên tệp hoặc thư mục)

	fmt.Println(filepath.IsAbs("dir/file"))  // trả về false vì đây là đường dẫn tương đố
	fmt.Println(filepath.IsAbs("/dir/file")) //trả về true vì đây là đường dẫn tuyệt đối (trên hệ thống Unix)

	filename := "config.json"
	ext := filepath.Ext(filename) //trả về phần mở rộng của tệp tin
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext)) //oại bỏ phần mở rộng khỏi tên tệp

	rel, err := filepath.Rel("a/b", "a/c/t/file") //trả về đường dẫn tương đối từ from đến to
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
}
