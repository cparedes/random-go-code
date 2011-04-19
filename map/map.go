package main

import "fmt"

func main() {
  m := map[int]string{1:"one" , 2:"two"}

  for _, v := range m {
    fmt.Println(v)
  }
}
