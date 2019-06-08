package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	//s3, s4, s5 = [5 6 10] [5 6 10 11] [5 6 10 11 12]
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	//arr = [0 1 2 3 4 5 6 10]
	fmt.Println("arr =", arr)

	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 = []int{2, 4, 6, 8}
	//[2 4 6 8], len=4, cap=4
	printSlice(s1)

	s2 = make([]int, 16)
	//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)

	s3 = make([]int, 10, 32)
	//[0 0 0 0 0 0 0 0 0 0], len=10, cap=31
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	//[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	//[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)


	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)

	printSlice(s2)
}
