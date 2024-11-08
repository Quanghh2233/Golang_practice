package main

import "fmt"

func Map[K, V any](s []K, transform func(K) V) []V {
	//K, V -> parameter, any đại diện cho bất kỳ dữ liệu nào
	//s là 1 slice có kiểu phần tử là K, đại diện cho các phần tử muốn chuyển đổi
	//transform -> 1 func nhận 1 tham số kiểu K và trả về 1 tham số kiểu V, là logic biến đổi có thể áp dụng cho từng phần tử mảng s
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	//dùng for lấy từng phần tử v trong s và áp dụng transform vào phần tử đó
	//kết quả của transform(v) đc lưu vào vị trí tương ứng i trong rs
	//mảng rs chứa các giá trị mới đc tạo ra từ mảng s qua transform và đc trả về
	return rs
}
func sliceIdx[S ~[]E, E comparable](s S, v E) int {
	// S ~[]E => S là 1 loại slice []E hoặc tương tự
	// E là kiểu của các phần tử trong slice, với ràng buộc là comparable(các phần tử của slice phải có khả năng so sánh == và !=) => cần thiết để kiểm tra v có bằng phần tử trong slice k
	// s S: 1 slice(hoặc tương tự) chứa các phần tử E
	// v E: giá trị muốn tìm trong s
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type SignedInt interface { //Constaint định nghĩa các kiểu dữ liệu mà Min sẽ chấp nhận
	int | int32 | int64 | int16 | int8
}

func Min[T SignedInt](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		if r > v {
			r = v
		}
	}
	return r
}
func main() {
	arr := []int{1, 2, 3}
	resultArr := Map(arr, func(v int) int { return v * 2 })
	//nhận arr làm mảng đầu vào []K, với K là int
	//transform = func(v int) int { return v * 2 } -> nhận só nguyên v và trả về v*2
	fmt.Println(resultArr)
	arr2 := []string{"a", "b", "c"}
	resultArr2 := Map(arr2, func(v string) string { return v + v + v })
	// K là string
	// nhận chuỗi v và trả về v+v+v
	fmt.Println(resultArr2)

	var s = []string{"foo", "bar", "zoo", "zodi"}
	fmt.Println("index of zoo:", sliceIdx(s, "zodi"))

	nums := []int{5, 3, 4, 8, 7, 6, 9, 2}
	Minnums := Min(nums)
	fmt.Println("Min:", Minnums)

	Ints := []int{1000, 200, 468, 78, 6126, 02056, 0, 498, 621, 232, -2, -16541, -62161, 98456, -78974}
	MinInts := Min(Ints)
	fmt.Println("Min:", MinInts)
}
