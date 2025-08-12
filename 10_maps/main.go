package main

import "fmt"

func main() {
	//creating map
	m := make(map[string]string)
	//setting elements
	m["name"] = "golang"
	fmt.Println(m["name"])
	m["name"] = "go"
	fmt.Println(m["name"])
	fmt.Println(len(m))
	m["type"] = "programming"
	fmt.Println(m)
	delete(m, "type")
	clear(m)
	fmt.Println(m["name"])

}
