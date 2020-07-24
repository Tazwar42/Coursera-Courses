package main

import (
    "fmt"
    "encoding/json"
)


func main() {
    var input1 string
    var input2 string
    m := make(map[string]string)

    fmt.Print("Enter Name: ")
    fmt.Scanln(&input1)
    fmt.Print("Enter Address: ")
    fmt.Scanln(&input2)


    m["name"] = input1
    m["address"] = input2

    //fmt.Println("Name:", m["name"])
    //fmt.Println("Address:", m["address"])
    jsonString, err := json.Marshal(m)
    if err != nil {
        panic(err)
    }

    //fmt.Println(datas)
    fmt.Println(string(jsonString))

}
