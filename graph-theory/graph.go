package main

import (
	"fmt"
)
func main() {
	graph := make(map[interface{}][]interface{})
	graph = map[interface{}][]interface{} { 
		"aaa" : {43,43,5} ,
	}
	graph["www"] = []interface{} {
		43,43,5,
	}
	fmt.Println(graph)
}