package main

import (
    "fmt"
)


var swapped = true

func swap(input []int, i int){
  input[i], input[i-1] = input[i-1], input[i]
}

func bubbleSort(input []int) {

    n:= 10
    swapped := true

    for swapped {
        swapped = false
        for i := 1; i < n; i++ {
                if input[i-1] > input[i] {
                    swap(input,i)
    		            swapped = true
                  }

            }
        }

    }



func main() {

    //fmt.Println("Enter the numbers of integer you want to sort")
    //fmt.Scan(&n)
    a := make([]int, 10)
    fmt.Println("Please enter the numbers.......:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Number (%d)----> ",i+1)
        fmt.Scan(&a[i])
      }

    fmt.Println("Given Numbers.........:")
    fmt.Println(a)
    bubbleSort(a)
    fmt.Println("Sorted numbers.............:")
    fmt.Println(a)


}
