package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	arr := []interface{} { 3, 2, 1, 4, "5", 6, 0, 3, 89, 1, 0, 8, 3,}
	quickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr);
}

func swap(arr []interface{}, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
} 

func quickSort(arr []interface{}, left, right int) {

	if left >= right {
		return
	}	

	pivot := arr[left];
	i := left + 1
	j := right

	for i <= j {
		if cvType(arr[i]) > cvType(pivot) && cvType(arr[j]) < cvType(pivot) {
			swap(arr, i, j)
		}
		if cvType(arr[i]) <= cvType(pivot) {
			i ++
		}
		if cvType(arr[j]) >= cvType(pivot) {
			j --
		}
	}

	swap(arr, j, left)
	quickSort(arr, left, j - 1)
	quickSort(arr, j + 1, right)
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