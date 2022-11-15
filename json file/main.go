package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Person struct {
    Name string
    Age string
    Gender int
}

func main() {

    bytes, err := ioutil.ReadFile("config.json")

    if err != nil {
        fmt.Println("Unable to load config.json file!")
        return
    }

    var myBook Person
    err = json.Unmarshal(bytes, &myBook)

    if err != nil {
        fmt.Println("JSON decode error!")
        return
    }

    fmt.Println(myBook) 
}