package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)

	fmt.Println("Traversing map")

	for k, _ := range m {
		fmt.Println(k)
	}

	fmt.Println("Getting values")
	if courseName, ok := m["name"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("Key does not exist")
	}
	fmt.Println("Deleting values")
	delete(m, "name")
	if courseName, ok := m["name"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("Key does not exist")
	}
}
