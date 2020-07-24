package main

import (
    "fmt"
    "os"
)


func GenDisplaceFn(a,v0,s0  float64) func (float64) float64{
   fn := func (t float64) float64 {
     		return ((.5*t*t*a)+(v0*t)+(s0))
         }
   return fn
}

func main() {


    var a float64
    fmt.Println("Enter  flaot value of a : ")

     _, err1 := fmt.Scanf("%f\n", &a)

    if err1 != nil {
      fmt.Println(err1)
    }

    fmt.Println("You have entered value of a: ", a)

    var v0 float64
    fmt.Println("Enter  flaot value of v0 : ")

     _, err2 := fmt.Scanf("%f\n", &v0)

    if err2 != nil {
      fmt.Println(err2)
    }

    fmt.Println("You have entered value of v0: ", v0)

    var s0 float64
    fmt.Println("Enter  flaot value of s0 : ")

     _, err3 := fmt.Scanf("%f\n", &s0)

    if err3 != nil {
      fmt.Println(err3)
    }

    fmt.Println("You have entered value of s0: ", s0)

    s1 := GenDisplaceFn(a,v0,s0)



    var t float64

    for {
    fmt.Println("Enter  flaot value of t (Time) or type exit to exit the program : ")

     _, err4 := fmt.Scanf("%f\n", &t)

    if err4 != nil {
      os.Exit(3)
    }

    fmt.Println("You have entered value of t (Time): ", t)


    fmt.Println("Displacement is: ",s1(t))
    }



}
