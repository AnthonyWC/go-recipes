package main

import (
	"fmt"
)

func main() {
	//	//	x := make([]int, 3, 5)
	//	//	x[0] = 10
	//	//	x[1] = 20
	//	//	x[2] = 30
	//	//	fmt.Println(x)

	//	y := make([]int, 3, 5)
	//	y[0] = 10
	//	y[1] = 20
	//	y[2] = 30
	//	fmt.Println(y)
	//	fmt.Println(len(y))
	//	fmt.Println(len(y))

	//	//	z := []int{10, 20, 30}
	//	//	fmt.Println(len(z))
	//	//	fmt.Println(len(z))

	//	z := []int{0: 10, 2: 30}
	//	fmt.Println(len(z))
	//	fmt.Println(len(z))

	x := []int{10, 20, 30}
	y := append(x, 40, 50)
	fmt.Println(x, y)

}
