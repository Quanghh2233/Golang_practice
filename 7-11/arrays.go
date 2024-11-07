package main

import "fmt"

func main() {
	var matx [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			matx[i][j] = i + j
		}
	}
	fmt.Println("2D: ", matx)
	matx = [3][4]int{
		{2, 3, 4},
		{1, 2, 4},
	}
	fmt.Println("2D: ", matx)
}
