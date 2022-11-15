package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
    Name string
    Age string
    Gender int
}

func main() {
    data := Person{"Hello Golang", "Roshan Jalan", 2022}
    bytes, _ := json.Marshal(data)
    fmt.Println(string(bytes)) // output:=  {"Title":"Hello Golang","Author":"Roshan Jalan","Year":2022}
}
