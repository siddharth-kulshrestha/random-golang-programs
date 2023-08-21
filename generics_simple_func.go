package main

import (
	"fmt"
)

type Num interface {
	int | float64
}

func Mapper[T Num](input []T, modifier func(t T) T) []T {
	output := []T{}
	for _, n := range input {
		output = append(output, modifier(n))
	}
	return output
}

func main() {
	fmt.Println(Mapper([]int{1, 2, 3, 4}, func(a int) int { return a * 2 }))
	fmt.Println(Mapper([]float64{1.2, 2.5, 3.6, 4.1}, func(a float64) float64 { return a * 2 }))
}
