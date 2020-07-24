package main

import (
        "fmt"


)

func main() {

       for {
        var f float64

        fmt.Printf("Enter a float value:")
        _, err := fmt.Scanf("%f\n", &f)

        var y int = int(f)
        if err != nil {
       break
    }

        fmt.Printf("You have entered:%d\n", y)
       }



}
