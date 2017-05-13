package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, world\n")

	var arr []int

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 4)
	arr = append(arr, 8)
	arr = append(arr, 10)

	var arr2 []int

	arr2 = append(arr2, 10)
	arr2 = append(arr2, 20)
	arr2 = append(arr2, 30)
	arr2 = append(arr2, 50)

	arr = append(arr, arr2...)

	println(len(arr))
	println(cap(arr))

	fmt.Println(arr)
}