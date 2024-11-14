package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f" //theo cấu trúc scheme://user:password@host:port/path?query#fragment.

	u, err := url.Parse(s) //chuyển chuỗi URL s thành một cấu trúc url.URL, giúp dễ dàng truy cập các thành phần khác nhau của URL.
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host) //chia thành tên host và port riêng.
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery) // phân tích chuỗi truy vấn k=v thành một map map[k:[v]].
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
