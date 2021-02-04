package main

import "fmt"

func main() {
  m := make(map[int]int)
  m[0] = 0
  changeMap(m)
  fmt.Println(m)
  fmt.Println("---")
  s := make([]int, 1)
  changeSlice(s)
  fmt.Printf("4s:%p\n", &s)
  appendSlice(s)
  fmt.Printf("2s:%p\n", &s)
}

func changeMap(m map[int]int) {
  m[0] = 1
}

func changeSlice(s []int) {
  fmt.Printf("5s:%p\n", &s)
  s[0] = 0
  fmt.Printf("3s:%p\n", &s)
}

func appendSlice(s []int) {
  fmt.Printf("6s:%p\n", &s)
  s = append(s, 1)
  fmt.Printf("1s:%p\n", &s)
}
