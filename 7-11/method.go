package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // có thể lấy receiver là con trỏ *rect
	return r.width * r.height
}
func (r *rect) changeStats(newW, newH int) { //=> method thay đổi chỉ số của đa giác với reciever là con trỏ
	r.width = newW
	r.height = newH
}
func (r rect) perim() int { // cũng có thể lấy receiver là giá trị rect
	return 2*r.width + 2*r.height
}
func (r rect) changeStat(newW, newH int) { //=> method thay đổi chỉ số của đa giác với reciever là giá trị
	r.width = newW
	r.height = newH
}
func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("Start")
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	fmt.Println("\nValue as receiver") //=> value sẽ không hiển thị sự thay đổi xảy ra trong method changeStat
	r.changeStat(10, 10)
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	fmt.Println("\nPointer as receiver") //=> pointer có thể truy cập vào giá trị tại địa chỉ rect => các thay đổi từ method ảnh hưởng đến kết quả được hiển thị
	r.changeStats(20, 10)
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())
}
