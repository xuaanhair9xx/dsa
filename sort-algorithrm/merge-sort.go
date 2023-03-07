package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int { 3, 2, 1, 4, 5, 6, 0, 3, 89, 1, 0}
	mergeSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}

func mergeSort(arr []int, left, right int) {
	if (left == right) {
		return
	}

	mid := (left + right) / 2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid + 1, right)
	merge(arr, left, mid, right)

}

func merge(arr []int, left, mid, right int) {
	var arrLeft []int
	var arrRight []int
	var arrTemp []int

	for i := left; i <= right ; i++ {
		if (i <= mid) {
			arrLeft = append(arrLeft, arr[i])
		} else {
			arrRight = append(arrRight, arr[i])
		}
	}
	arrLeft = append(arrLeft, math.MaxInt)
	arrRight = append(arrRight, math.MaxInt)

	size := right - left + 1
	
	k := 0 
	i := 0
	j := 0

	for (k < size) {
		if (arrLeft[i] < arrRight[j]) {
			arrTemp = append(arrTemp, arrLeft[i])
			i++
		} else {
			arrTemp = append(arrTemp, arrRight[j])
			j ++
		}
		k++
	}

	k = 0 
	for i := left; i <= right; i++ {
		arr[i] = arrTemp[k]
		k++
	} 
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