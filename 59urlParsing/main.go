package main

import (
	"fmt"
	"net/url"
	"net"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)

	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		panic(err)
	}
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	fmt.Println(m["k"])
	fmt.Println(m["k"][0])
}