package main

import "fmt"

func main() {
	nums := []int{2, 6, 8, 1}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums { // => xuất vị trí đứng của giá trị trong mảng
		if num == 6 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "pineapple", "c": "starfruit"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for _, v := range kvs { //=> _ dùng để đại diện cho các giá trị ko dùng đến(ở đây là key)
		fmt.Println("value:", v)
	}
	fmt.Println("len:", len(kvs))
	for i, c := range "Quang" { //=> i đại diện cho chỉ số của ký tự(tương tự như trong mảng), c đại diện cho mã unicode của ký tự
		fmt.Println(i, c, string(c))

	}
}
