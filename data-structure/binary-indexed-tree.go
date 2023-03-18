package main

import(
	"fmt"
)

var arr []int
var m_array []int
func main() {
	arr = []int{ 1, 7, 3, 0, 5, 8, 3, 2, 6, 2, 1, 1, 4, 5 };
	m_array = make([]int, len(arr) + 1)
	Init()
	fmt.Println(m_array)
	fmt.Println(prefix_query(12))
	update(4, 2)
	fmt.Println(m_array)
	fmt.Println(prefix_query(12))
}

func Init() {
	for i := 0; i < len(arr); i++ {
		m_array[i+1] = arr[i]
	}
	for i := 1; i < len(m_array); i++ {
		i2 := i + (i & -i);
		if i2 < len(m_array) {
			m_array[i2] += m_array[i]
		}
	}
}

func prefix_query(idx int) int {
	result := 0
	for i := idx+1; i > 0; i = i - (i & -i) {
		result = result + m_array[i]
	}
	return result
}

func range_query(from, to int) {
	if from == 0 {
		return prefix_query(to)
	} else {
		return prefix_query(to) - prefix_query(from - 1)
	}
}

func update(idx, add int) {
	for i := idx+1; i < len(m_array); i = i + (i & -i) {
		m_array[i] = m_array[i] + add
	}
}