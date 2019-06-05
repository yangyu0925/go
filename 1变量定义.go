package main

import "fmt"

//var (
//	aa = 33
//	ss = "kkk"
//)
//
//func variable()  {
//	var a int
//	var s string
//	fmt.Println(a, s)
//	fmt.Printf("%d %q \n", a, s)
//}
//
//func variableInit()  {
//	var a, b int = 3, 4
//	var s string = "abc"
//	fmt.Println(a, b, s)
//}
//
//func variableType()  {
//	var a, b, c, s = 3, 4, true, "hello"
//	fmt.Println(a, b, c, s)
//}
//
//func variableShort()  {
//	a, b, c, s := 3, 4, true, "hello world"
//	fmt.Println(a, b, c, s)
//}

//func main() {
//	fmt.Println("hello world")
//	variable()
//	variableInit()
//	variableType()
//	variableShort()
//	a, b, i, s1, s2 := true, false, 3, "hello", "world"
//	fmt.Println(a, b, i, s1, s2)
//}

var (
	name = "yangyu"
	age = "33"
)

func variable()  {
	var a, b, c, d = 1, 2, true, "yang"
	fmt.Println(a, b, c, d)
}

func variableShrt()  {
	a, b, c, d := 1, 2, 3, true
	fmt.Println(a, b, c, d)
}

func main()  {
	variableShrt()
}
