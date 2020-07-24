package main

import (
    "time"
    "fmt"
)

var x int

func sum() {
    x++
    fmt.Println(x)
}

func sub() {
    x--
    fmt.Println(x)
}

func main() {
    go sum()
    go sub()
    time.Sleep(1 * time.Second)
}

/*
Race Conditions Outcome depends on non-deterministic ordering
Races occur due to communication
In this example communication depends on x which is the shared variable and thats where Race
Conditions occur.
If there was no shared variable there would be never race Condition becasue ordering would
be different.
So communication is the source of race Condition and its very common in goroutine

*/
