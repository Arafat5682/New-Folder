package main

import "fmt"

func main() {
   var students[3]string
  students[0] = "Atiq"
  students[1] = "Aslam"
  students[2] = "Arif"
x := students[0:3]
fmt.Println(x)
}