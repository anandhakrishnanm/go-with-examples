package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffe"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, err := xml.MarshalIndent(coffee, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	tomato := &Plant{Id: 81, Name: "Tomatto"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, err = xml.MarshalIndent(nesting, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
