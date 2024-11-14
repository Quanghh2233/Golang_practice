package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())      //trả về Unix timestamp tính bằng giây kể từ ngày 1 tháng 1 năm 1970 (Epoch time). Kết quả là một số nguyên thể hiện số giây.
	fmt.Println(now.UnixMilli()) //trả về Unix timestamp tính bằng milliseconds (ms) kể từ Epoch
	fmt.Println(now.UnixNano())  // trả về Unix timestamp tính bằng nanoseconds (ns) kể từ Epoch.

	fmt.Println(time.Unix(now.Unix(), 0))     //chuyển đổi Unix timestamp (số giây) thành đối tượng time.Time. Tham số đầu tiên là số giây, và tham số thứ hai là số nanoseconds.
	fmt.Println(time.Unix(0, now.UnixNano())) //chuyển đổi Unix timestamp nanoseconds thành time.Time. Trong trường hợp này, tham số đầu tiên là số giây (bằng 0), và tham số thứ hai là số nanoseconds.
}
