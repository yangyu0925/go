package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	//m2 == empty map
	m2 := make(map[string]int)

	//m3 == nil
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m{
		fmt.Println(k,":",v)
	}

	fmt.Println("Getting value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if coursName, ok := m["cours"]; ok{
		fmt.Println(coursName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	delete(m, "name")
	fmt.Println(name, ok)

	name, ok = m["name"]
	fmt.Println(name, ok)
}
