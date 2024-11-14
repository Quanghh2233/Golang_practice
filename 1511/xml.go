package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"` //tên của phần tử xml chính <plant> khi chuyển đổi struct thành xml
	Id      int      `xml:"id"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out)) // xml.Header là một hằng số trong gói xml chứa chuỗi <?xml version="1.0" encoding="UTF-8"?>.
	// => thêm phần đầu này vào trước xml của coffee, tạo ra 1 chuỗi xml hoàn chỉnh
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil { //chuyển đổi XML ngược lại thành 1 đối tượng plant
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	potato := &Plant{Id: 88, Name: "Potato"}
	potato.Origin = []string{"Spain", "London"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"` // mỗi plant sẽ nằm trong <parent><child><plant>
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato, potato} // tạo nesting với Plants là mảng chứa coffee, tomato
	out, _ = xml.MarshalIndent(nesting, " ", " ")     // Chuyển đổi nesting thành XML với cấu trúc được lồng ghép. Kết quả được in ra.
	fmt.Println(string(out))

}
