package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Mumbai Indians"
	strEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(strEnc)

	strDec, err := base64.StdEncoding.DecodeString(strEnc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(strDec))
	fmt.Println()

	urlEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEnc)
	urlDec, err := base64.URLEncoding.DecodeString(urlEnc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(urlDec))
}