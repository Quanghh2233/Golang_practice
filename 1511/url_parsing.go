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
	fmt.Println("-------------------------------------------------")
	p1 := fmt.Println
	s2 := "Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"
	u1, err1 := url.Parse(s2)
	if err1 != nil {
		panic(err1)
	}
	p1(u1.Scheme)
	p1(u1.User)
	p1(u1.User.Username())
	pass, _ := u1.User.Password()
	p1(pass)
	p1(u1.Host)
	p1(u1.Port)
	p1(u1.Path)
	p1(u1.Fragment)
	p1(u1.RawQuery)
	k, _ := url.ParseQuery(u1.RawQuery)
	p1(k)
	p1(k["key2"][0])
}
