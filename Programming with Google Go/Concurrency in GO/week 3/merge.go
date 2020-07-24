package main

import (
  "fmt"
  "sort"
  "sync"
  "os"

)

func array1(slice1 []int, wg1 *sync.WaitGroup){
    fmt.Println("Unsorted Numbers of 1st Goroutine:",slice1)
    sort.Ints(slice1)
    wg1.Done()
}

func array2(slice2 []int,  wg2 *sync.WaitGroup){
    fmt.Println("Unsorted Numbers of 2nd Goroutine:",slice2)
    sort.Ints(slice2)
    wg2.Done()
}

func array3(slice3 []int,  wg3 *sync.WaitGroup){
    fmt.Println("Unsorted Numbers of 3rd Goroutine:",slice3)
    sort.Ints(slice3)
    wg3.Done()
}

func array4(slice4 []int,  wg4 *sync.WaitGroup){
    fmt.Println("Unsorted Numbers of 4th Goroutine:",slice4)
    sort.Ints(slice4)
    wg4.Done()
}



func main() {

  fmt.Println("Please give the number of integers(which can be divided by 4) to sort...suppose {4,8,12,16,20,24}.....etc")

  var n int
  fmt.Scan(&n)
  a := make([]int, n)
  for i := 0; i < n; i++ {
      fmt.Printf("Number %d--->",i)
      fmt.Scan(&a[i])
    }

  var num = len(a)

  if num%4 != 0 {
    fmt.Println("invalid numbers of integers given. Try Again.Please give the number of integers(which can be divided by 4) to sort...suppose {4,8,12,16,20,24}.....etc")
    os.Exit(3)
  }

  partition := int(num / 4)
  var (
        s1 = make([]int, partition)
        s2 = make([]int, partition)
        s3 = make([]int, partition)
        s4 = make([]int, partition)
    )

  for i := 0; i < num; i++ {
      if i < partition {
            s1[i] = a[i]
        } else if i< partition+partition {
            s2[i-partition] = a[i]
        } else if i < partition +partition +partition {
            s3[i-partition-partition] = a[i]
        } else {
            s4[i-partition-partition-partition] = a[i]
        }
    }



  var wg1 sync.WaitGroup
  wg1.Add(1)
  go array1(s1, &wg1)
  wg1.Wait()

  var wg2 sync.WaitGroup
  wg2.Add(1)
  go array2(s2, &wg2)
  wg2.Wait()

  var wg3 sync.WaitGroup
  wg3.Add(1)
  go array3(s3, &wg3)
  wg3.Wait()

  var wg4 sync.WaitGroup
  wg4.Add(1)
  go array4(s4, &wg4)
  wg4.Wait()

  merge1 := merge(s1,s2)
  merge2 := merge(s3,s4)
  merge_total := merge(merge1,merge2)
  fmt.Println("Sorted Interges are.....:",merge_total)

}


func merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))

    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }

    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }

    return
}
