package main

import "fmt"

func swap(a, b *int)  {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int)  {
	a, b = b, a;
	return a, b;
}

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	a, b = swap1(a, b)
	fmt.Println(a, b)
}
