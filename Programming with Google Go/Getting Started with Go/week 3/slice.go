package main

import (
        "fmt"
        "sort"
)

func main() {
  s := make([]int, 0, 3)
  for{
  var i int
  fmt.Printf("Enter a integer value:")
  _, err := fmt.Scanf("%d\n", &i)
  if err != nil {
  break
  }
  s = append(s, i)
  sort.Ints(s)
  fmt.Println(s)
}
}
