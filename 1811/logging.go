package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger") //Dòng log cơ bản sử dụng logger mặc định của Go, kèm theo thời gian.

	log.SetFlags(log.LstdFlags | log.Lmicroseconds) //Thêm microseconds vào thông tin thời gian.
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile) //Thêm thông tin file và dòng hiện tại vào log.
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags) //Tạo logger mới với tiền tố "my:".
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:") //Đổi tiền tố thành "ohmy:"
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags) //Ghi log vào bộ nhớ tạm (buf) thay vì ghi ra console.

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil) //Tạo một logger để ghi log dưới dạng JSON
	myslog := slog.New(jsonHandler)
	myslog.Info("hello there")

	myslog.Info("hello again", "key", "val", "age", 25)

}
