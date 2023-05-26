package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page    int
	Fruites []string
}

type response2 struct {
	Page    int      `json:"page"`
	Fruites []string `json:"fruites"`
}

func main() {
	bolB, err := json.Marshal(true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bolB))

	intB, err := json.Marshal(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(intB))

	fltB, err := json.Marshal(2.34)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(fltB))

	strB, err := json.Marshal("gopher")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, err := json.Marshal(slcD)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 10, "banana": 30}
	mapB, err := json.Marshal(mapD)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(mapB))

	res1D := response1{
		Page:    1,
		Fruites: []string{"apple", "peach", "pear"},
	}
	res1B, err := json.Marshal(res1D)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res1B))

	res2D := response2{
		Page:    1,
		Fruites: []string{"apple", "peach", "pear"},
	}
	res2B, err := json.Marshal(res2D)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res2B))

	byt := []byte(`{"num": 10.1234, "strs": ["a", "b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 2, "fruites": ["apple", "peach"]}`
	res := response2{}
	err = json.Unmarshal([]byte(str), &res)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println(res.Fruites[0])
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
