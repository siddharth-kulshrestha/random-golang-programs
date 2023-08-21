package main

import (
  "fmt"
)


func Add(a int, b int) int {
  return a + b
}

func AddFloat(a float64, b float64) float64 {
  return a + b
}

func AddGeneric[T float64 | int](a T, b T) T {
 return a + b
}


func main() {
  fmt.Println(Add(1, 2))
  fmt.Println(AddGeneric(1.2, 1.6))
  fmt.Println(AddGeneric(1, 2))
}
