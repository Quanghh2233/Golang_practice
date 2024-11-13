package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type resp1 struct {
	Page   int
	Fruits []string
}

type resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//ử dụng json.Marshal để chuyển đổi một số kiểu dữ liệu như bool, int,
	//float64, và string thành định dạng JSON.
	//Kết quả là các giá trị này được in ra dưới dạng chuỗi JSON.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//chuyển đổi một slice ([]string) sang chuỗi JSON.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	//chuyển đổi một map (map[string]int) sang chuỗi JSON.
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//Định nghĩa cấu trúc resp1 và resp2, sau đó chuyển đổi các đối tượng của cấu trúc này sang JSON.
	res1D := &resp1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	//Lưu ý rằng resp2 sử dụng thẻ JSON (json:"...") để đặt tên khóa tùy chỉnh trong chuỗi JSON.
	res2D := &resp2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13, "str":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil { //json.Unmarshal được sử dụng để giải mã một chuỗi JSON vào một map (map[string]interface{}).
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) //ép kiểu
	fmt.Println(num)

	strs := dat["str"].([]interface{}) //ép kiểu
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits":["apple","peach"]}`
	res := resp2{}
	json.Unmarshal([]byte(str), &res) //chuyển chuỗi JSON thành một đối tượng của cấu trúc resp2 và truy cập phần tử đầu tiên trong mảng Fruits
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout) //json.NewEncoder tạo một encoder
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) //ghi trực tiếp JSON vào os.Stdout
}
