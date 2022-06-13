package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` // Use "name" instead of "Name" for the JSON
	Age  int    `json:"age"`  // field name.
}

func main() {
	p1 := Person{"Jake", 39}
	p1String, err := json.MarshalIndent(p1, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("struct -> JSON string: ", string(p1String))

	p2String := []byte(`{"name":"Kristin", "age": 38}`)
	var p2 Person
	if err := json.Unmarshal(p2String, &p2); err != nil {
		fmt.Println(err)
		return
	}
	//print the struct
	fmt.Println("JSON string -> struct", p2)
}
