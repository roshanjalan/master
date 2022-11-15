package main

import (
    "fmt"
    "encoding/json"
)

type Book struct {
    Title string `json:"title"`
    Author string `json:"author"`
    Year int `json:"year"`
}

func main() {
    myBook := `{
        "title": "Hello Golang",
        "author": "Roshan Jalan",
        "year": 2022
    }`

    var Books Book
    err:= json.Unmarshal([]byte(myBook), &Books)

    if err != nil {
        fmt.Println("JSON decode error!")
        return 
    }

    fmt.Println(Books) // {Hello Golang Roshan Jalan 2022}
}
