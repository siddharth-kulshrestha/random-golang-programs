package main

import (
	"fmt"
)

type MyInt int

type Num interface {
	int | float64
}

type Num2 interface {
	~int | float64
}

func AddGeneric[T Num](a T, b T) T {
	return a + b
}

func AddGeneric2[T Num2](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(AddGeneric(1.2, 1.6))
	fmt.Println(AddGeneric(1, 2))
	fmt.Println(AddGeneric2(MyInt(5), MyInt(6)))
	fmt.Println(AddGeneric2(MyInt(5), 2))
	// fmt.Println(AddGeneric2(MyInt(5), 2.8)) // Can't use it like this.
	fmt.Println(AddGeneric2(5, 2))
	fmt.Println(AddGeneric2(5.8, 2.9))
	fmt.Println(AddGeneric2(5, 2.9))

}
