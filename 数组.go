package main

import "fmt"

func printArray(arr [5]int)  {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3{
		fmt.Println(i, v)
	}

	printArray(arr1)
}
