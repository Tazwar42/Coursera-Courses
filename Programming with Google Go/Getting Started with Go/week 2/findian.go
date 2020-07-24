package main

import (
        "fmt"
        "regexp"


)

func main() {

    var str string
    fmt.Println("Please enter the string:")
    fmt.Scanf("%s", &str)

    match, _ := regexp.MatchString("(?i)^i.*a+.*(?i)n$", str)
    if match == true{
      fmt.Println("Found!!")
    } else {
      fmt.Println("Not found")
    }



}
