package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div1(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apple(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int{
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(eval(3, 4, "x"))
	fmt.Println(div(13, 3))
	fmt.Println(div1(13, 3))
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apple(pow, 3, 4))
	fmt.Println(apple(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	},3, 4))

	fmt.Println(sum(1, 2, 3))
}
