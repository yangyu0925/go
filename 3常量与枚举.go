package main

import "fmt"

func enums()  {

	const a = 3

	//const(
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)

	const(
		cpp = iota
		java
		python
		golang
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(a)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	enums()
}
