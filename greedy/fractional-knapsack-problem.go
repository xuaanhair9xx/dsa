package main

import(
	"fmt"
)

type Item struct {
	value int
	weight int
}

func main()  {
	arr := []Item {
		Item { value: 500, weight: 30 }, 
		//Item { value: 100, weight: 20 },
		//Item { value: 120, weight: 30 },
	}
	w := 10
	knapsack(arr, w)
}

func knapsack(items []Item, w int) {
	bubbleSort(items)
	var finalResult float64 = 0
	for _, item := range items {
		if item.weight < w {
			finalResult = finalResult + float64(item.value)
			w = w - item.weight
		} else {
			finalResult = finalResult + float64(item.value) / float64(item.weight) * float64(w)
			break
		}
	}
	fmt.Println(finalResult)
}

func cmp(a Item, b Item) bool {
	v1 := float64(a.value) / float64(a.weight)
	v2 := float64(b.value) / float64(b.weight)
	if v1 > v2 {
		return true
	}
	return false
}

func swap(arr []Item, x, y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

func bubbleSort(arr []Item) {
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		for j:= 0; j < n - 1 - 1; j++ {
			if cmp(arr[j], arr[j+1]) {
				swap(arr, j, j+ 1)
			}
		}
	}
}

