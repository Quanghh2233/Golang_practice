package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8092", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	//w (http.ResponseWriter): Đối tượng ghi dữ liệu phản hồi (response) về cho client.
	//req (*http.Request): Đối tượng chứa thông tin của yêu cầu (request) từ client.

}
func headers(w http.ResponseWriter, req *http.Request) { //trả về một map với khóa là tên header và giá trị là danh sách các giá trị header
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
