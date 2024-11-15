package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //Tạo một Scanner để đọc từ đầu vào chuẩn

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text()) // Đọc từng dòng từ đầu vào chuẩn cho đến khi gặp kết thúc của đầu vào
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
