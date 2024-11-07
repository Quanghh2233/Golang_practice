package main

import (
	"fmt"
	"slices"
)

func main() {
	var s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// fmt.Println(s)
	// fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[1:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"q", "h", "h"}
	fmt.Println("dcl:", t)

	t2 := []string{"q", "h", "h"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	matx := make([][]int, 4)
	for i := 0; i < 4; i++ {
		innerLen := i + 1
		matx[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matx[i][j] = i + j
		}
	}
	fmt.Println("2D:", matx)
}
