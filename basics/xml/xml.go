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

// to get Plant struct in pritable format
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// first is prefix: single space
	// second is indent: double space
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))
	// <plant id="27">
	//   <name>Coffee</name>
	//   <origin>Ethiopia</origin>
	//   <origin>Brazil</origin>
	// </plant>

	// We can understand this better as
	fmt.Println()
	out2, _ := xml.MarshalIndent(coffee, "+", "--")
	fmt.Println(string(out2))
	// +<plant id="27">
	// +--<name>Coffee</name>
	// +--<origin>Ethiopia</origin>
	// +--<origin>Brazil</origin>
	// +</plant>
	fmt.Println()

	fmt.Println(xml.Header + string(out))
	// <?xml version="1.0" encoding="UTF-8"?>
	// <plant id="27">
	// 	...
	// </plant>

	// convert xml string to xml struct
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// nesting of xml in another xml
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// defining a struct inside a function
	// just like func inside func or class inside class in python

	// The parent>child>plant field tag tells the encoder to nest
	// all plants under <parent><child>...
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	// using *Plant instead of Plant because coffee and tomato are
	// pointer to struct i.e. *coffee and *tomato are actual value
	// or we can put Plants  []Plant `xml:"parent>child>plant"`
	// but then we have to append obj like
	// nesting.Plants = []Plant{*coffee, *tomato}
	// point is in Plant here is array of pointer to Plant struct
	// otherwise we can keep it as array of Plant struct, both are valid

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	fmt.Println()
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
	// 	<nesting>
	// 	<parent>
	// 	  <child>
	// 		<plant id="27">
	// 		  <name>Coffee</name>
	// 		  <origin>Ethiopia</origin>
	// 		  <origin>Brazil</origin>
	// 		</plant>
	// 		<plant id="81">
	// 		  <name>Tomato</name>
	// 		  <origin>Mexico</origin>
	// 		  <origin>California</origin>
	// 		</plant>
	// 	  </child>
	// 	</parent>
	//  </nesting>
}
