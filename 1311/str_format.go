package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p) //in giá trị của p dưới dạng {x y}

	fmt.Printf("struct2: %+v\n", p) //hiển thị chi tiết giá trị của p dưới dạng {x:1 y:2}

	fmt.Printf("struct3: %#v\n", p) // in ra struct dưới dạng khai báo đầy đủ, bao gồm cả package và struct

	fmt.Printf("type: %T\n", p) //in ra kiểu dữ liệu

	fmt.Printf("bool: %t\n", true) // in ra bool

	fmt.Printf("int: %d\n", 123) //số nguyên

	fmt.Printf("bin: %b\n", 14) // nhị phân

	fmt.Printf("char: %c\n", 33) //ký tự ASCII của 1 số nguyên

	fmt.Printf("hex: %x\n", 456) //hexadecimal

	fmt.Printf("float1: %f\n", 78.9) //in ra số thực với định dạng dấu phẩy động

	fmt.Printf("float2: %e\n", 123400000.0) //in ra số thực dưới dạng khoa học
	fmt.Printf("float3: %E\n", 123400000.0) //in ra số thực dưới dạng khoa học

	fmt.Printf("str1: %s\n", "\"string\"") //in chuỗi

	fmt.Printf("str2: %q\n", "\"string\"") //in chuỗi với dấu ""

	fmt.Printf("str3: %x\n", "hex this") //in chuỗi dưới dạng hexa

	fmt.Printf("pointer: %p\n", &p) //địa chỉ của bộ nhớ p

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) //số nguyên với độ rộng 6 ký tự

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) //số nguyên với độ rộng 6 ký tự và 2 chứ số thập phân

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) //số nguyên với độ rộng 6 ký tự, 2 chữ số thập phân và căn trái

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // chuỗi với độ rộng 6 ký tự

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b") // chuỗi với độ rộng 6 ký tự và căn trái

	s := fmt.Sprintf("sprintf: a %s", "string") //tạo 1 chuỗi định dạng mà k in ra, lưu vào biến
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") //in ra io.writer cụ thể(ví dụ os.stderr)
}
