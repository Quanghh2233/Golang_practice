package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") // kiểm tra chuỗi "peach" có khớp với biểu thức chính quy "p([a-z]+)ch" không. [a-z]+ đại diện cho 1 hoặc nhiều ký tự thường
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") //Compile tạo ra một đối tượng biểu thức chính quy r, có thể được sử dụng nhiều lần sau đó để tìm kiếm hoặc thay thế.

	fmt.Println(r.MatchString("peach"))      // khớp chuỗi
	fmt.Println(r.FindString("peach punch")) // tìm chuỗi đầu tiên khớp

	fmt.Println("ind:", r.FindStringIndex("peach punch")) // tìm chỉ mục của chuỗi đầu tiên khớp

	fmt.Println(r.FindStringSubmatch("peach punch")) // Tìm chuỗi con khớp

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // tìm chỉ mục của các chuỗi con khớp

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // tìm tất cả chuỗi khớp

	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //Tìm tất cả chỉ mục của chuỗi con khớp

	fmt.Println(r.FindAllString("peach punch pinch", 2)) //Tìm giới hạn số lượng chuỗi khớp

	fmt.Println(r.Match([]byte("peach"))) //Kiểm tra khớp chuỗi với byte slice

	r = regexp.MustCompile("p([a-z]+)ch") //MustCompile là phiên bản của Compile nhưng sẽ gây panic nếu biểu thức không hợp lệ.
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //Thay thế chuỗi khớp

	in := []byte("a peach")                    //Thay thế bằng hàm
	out := r.ReplaceAllFunc(in, bytes.ToUpper) // thay thế chuỗi khớp bằng kết quả của hàm bytes.ToUpper, biến chuỗi con khớp thành chữ hoa.
	fmt.Println(string(out))
}
