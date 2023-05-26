package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, err := strconv.ParseFloat("1.234", 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)

	i, err := strconv.ParseInt("123", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	d, err := strconv.ParseInt("0x1c8", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	u, err := strconv.ParseUint("789", 0, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	k, err := strconv.Atoi("135")
	if err != nil {
		panic(err)
	}
	fmt.Println(k)

	_, err = strconv.Atoi("wat")
	fmt.Println(err)

}
