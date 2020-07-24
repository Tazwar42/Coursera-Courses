package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"


)

type Name struct {
    fname string
    lname string
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    var file_name string
    fmt.Println("Please Enter the file name........:")
    fmt.Scanf("%s", &file_name)

    readFile, err := os.Open(file_name)

	  if err != nil {
		log.Fatalf("failed to open file: %s", err)
	  }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileTextLines []string
    for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	  }

    readFile.Close()
    names := []*Name{}


	  for _, eachline := range fileTextLines {
		s := strings.Split(eachline, " ")
    name1 := new(Name)
    name1.fname = s[0]
    name1.lname = s[1]
    names = append(names, name1)
	  }


    for i := range(names) {
        name := names[i]
        fmt.Printf("First Name: %s, Last Name: %s\n", name.fname,name.lname)

    }

}
