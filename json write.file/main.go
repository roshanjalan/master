package main

import (
    "io/ioutil"
    "encoding/json"
)

type Person struct {
    Name string
    Age string
    Gender int
}

func main() {
    data := Person{"Hello Golang", "Roshan Jalan", 2022}
    bytes, _ := json.Marshal(data)
	ioutil.WriteFile("config.json", bytes, 0644)
}

