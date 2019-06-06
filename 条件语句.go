package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return  result;
}

func grade(score int) string  {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "D"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}

	if contents1, err1 := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("%s \n", contents1)
	}

	fmt.Println(eval(9, 3, "/"))

	score := grade(110)
	fmt.Println(score)
}
