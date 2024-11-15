package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	err := os.Mkdir("subdir", 0755) //tạo một thư mục mới có tên subdir với quyền truy cập 0755
	//(quyền đọc, ghi, và thực thi cho chủ sở hữu, quyền đọc và thực thi cho nhóm và người khác)
	check(err)
	defer os.RemoveAll("subdir") // đảm bảo rằng thư mục subdir và tất cả các tệp con của nó sẽ được xóa khi hàm main() kết thúc

	createEmptyFile := func(name string) { //tạo ra một tệp trống với tên được truyền vào trong tham số name
		d := []byte("")
		check(os.WriteFile(name, d, 0644)) //tạo tệp và ghi một mảng byte rỗng vào tệp đó với quyền truy cập 0644
	}

	createEmptyFile("subdir/f1")

	err = os.MkdirAll("subdir/parent/child", 0755) //tạo toàn bộ các thư mục con cần thiết trong đường dẫn "subdir/parent/child", nếu chúng chưa tồn tại
	check(err)

	createEmptyFile("subdir/parent/f2")
	createEmptyFile("subdir/parent/f3")
	createEmptyFile("subdir/parent/child/f4")

	c, err := os.ReadDir("subdir/parent") //đọc danh sách các tệp và thư mục trong thư mục "subdir/parent"
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	err = os.Chdir("subdir/parent/child") //thay đổi thư mục làm việc hiện tại thành thư mục(tương tự cd)
	check(err)

	c, err = os.ReadDir(".") //đọc các tệp trong thư mục hiện tại (thư mục child), và sau đó danh sách các tệp được in ra.
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..") //quay lại thư mục gốc của chương trình.
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit) //đi qua tất cả các thư mục và tệp trong thư mục "subdir" và gọi hàm visit cho mỗi mục trong đó
}
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
