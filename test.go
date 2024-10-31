package main

import "fmt"

func prod(a []int) {
	a = append(a, 1)
	a[0] = 5
}

func main() {
	arr := make([]int, 5, 6)
	fmt.Println("Initial array:", arr)
	prod(arr)
	fmt.Println("Modified array:", arr)

	arr2 := make([]int, 5, 5)
	fmt.Println("Initial array:", arr2)
	prod(arr2)
	fmt.Println("Modified array:", arr2)
}
