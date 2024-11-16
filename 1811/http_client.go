package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://stackoverflow.com/") // gửi yêu cầu http get đến url chỉ định
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //Đảm bảo đóng response body sau khi sử dụng xong, để tránh rò rỉ tài nguyên.

	fmt.Println("Response status:", resp.Status) //trả về trạng thái HTTP của phản hồi, ví dụ: "200 OK", "404 Not Found"

	scanner := bufio.NewScanner(resp.Body)     //Tạo một scanner để đọc dữ liệu từ body của phản hồi từng dòng.
	for i := 0; scanner.Scan() && i < 5; i++ { //loop tối đa 5 dòng, kiểm tra xem có dòng tiếp theo không.
		fmt.Println(scanner.Text()) //lấy nội dung của dòng hiện tại.
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	resp2, er2 := http.Get("https://gemini.google.com/app")
	if er2 != nil {
		panic(er2)
	}
	defer resp2.Body.Close()

	fmt.Println("Response Status:", resp2.Status)

	reader := bufio.NewReader(resp2.Body)
	for i := 0; i < 6; i++ {
		line, err := reader.ReadString('\n') // Đọc đến ký tự xuống dòng
		if err != nil {
			break // Thoát nếu hết dữ liệu hoặc có lỗi
		}
		fmt.Println(line)
	}
}
