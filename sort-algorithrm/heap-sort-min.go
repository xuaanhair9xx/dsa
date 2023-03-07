package main

import (
	"fmt"
	"strconv"
	"log"
)

// Left node: 2n + 1
// Right node: 2n + 2
// Parent node: (n-1) / 2
// Last parent node: size/2 - 1
func main() {
	arr := []interface{} { 5, "7", 1, 8, 6, 9, 4, 2, 3}
	lastParentNode := len(arr) / 2 - 1
	buildMaxHeap(arr, lastParentNode) 
	printHeap(arr)
	
}

func printHeap(arr []interface{}) {
	fmt.Println(arr)
	for len(arr) > 0 {
		swap(arr, 0, len(arr)-1)
		fmt.Println(arr[len(arr)-1])
		arr = arr[:len(arr) - 1]
		rebuildHeapMax(arr, 0)
	}
}

func rebuildHeapMax(arr []interface{}, parentNode int) {
	leftNode := 2*parentNode + 1
	rightNode := 2*parentNode + 2
	largest := parentNode
	n := len(arr)

	if leftNode < n && cvType(arr[largest]) > cvType(arr[leftNode]) {
		largest = leftNode
	}

	if rightNode < n && cvType(arr[largest]) > cvType(arr[rightNode]) {
		largest = rightNode
	}

	if largest != parentNode {
		swap(arr, parentNode, largest)
		rebuildHeapMax(arr, largest)
	}

}

func swap(arr []interface{}, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
} 

func buildMaxHeap(arr []interface{}, i int) {
	if (i < 0) {
		return
	}
	leftNode := 2*i + 1
	rightNode := 2*i + 2
	largest := i

	if cvType(arr[largest]) > cvType(arr[leftNode]) {
		largest = leftNode
	}

	if cvType(arr[largest]) > cvType(arr[rightNode]) {
		largest = rightNode
	}

	if largest != i {
		swap(arr, i, largest)
		rebuildHeapMax(arr, largest)
	}
	buildMaxHeap(arr, i - 1)
}

func cvType(i interface{}) int {
	switch v := i.(type) {
		case int:
			return v
		case string:
			vC, err := strconv.Atoi(v)
			if err != nil {
				log.Panic(err)
				return 0
			}
			return vC
		default:
			log.Panic("Type is invalid.")
			return 0
	}
}