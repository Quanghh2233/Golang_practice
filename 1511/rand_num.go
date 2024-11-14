package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// in hai số nguyên ngẫu nhiên trong khoảng 0 đến 99, cách nhau bởi dấu phẩy.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	//in một số thực ngẫu nhiên, tiếp theo là hai số thực ngẫu nhiên trong khoảng từ 0 đến 1.0
	fmt.Println(rand.Float64())

	//in một số thực ngẫu nhiên, tiếp theo là hai số thực ngẫu nhiên trong khoảng từ 5.0 đến 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s2 := rand.NewPCG(42, 1024) //tạo một nguồn ngẫu nhiên mới với hạt giống là các giá trị 42 và 1024.
	r2 := rand.New(s2)          // tạo một rand.Rand mới dựa trên nguồn s2
	//sẽ tạo các số nguyên ngẫu nhiên trong khoảng từ 0 đến 99 từ nguồn r2
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	//sẽ tạo các số nguyên ngẫu nhiên trong khoảng từ 0 đến 99 từ nguồn r3
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
