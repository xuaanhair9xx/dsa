package main

import(
	"fmt"
	"math"
)

func main() {
	arr := []float64{1, 3, 2, 5, 4, 6, 7, 8, 9,10}
	n := len(arr)
	tree := make([]float64, (len(arr)+4)*2)
	treeIter := make([]float64, (len(arr)+4)*2)
	var buildTree func(int, int, int) 
	var buildTreeIter func()
	var find func(int, int, int, int, int) float64
	buildTree = func(index int, l, r int) {
		if l == r {
			tree[index] = arr[l]
			return
		}
		mid := l + (r - l) / 2
		buildTree(index*2 + 1, l, mid)
		buildTree(index*2 + 2, mid + 1, r)
		tree[index] =  math.Max(tree[index*2 + 1], tree[index*2+2])
	}

	buildTreeIter = func() {
		for i := 0; i < n; i++ {
			treeIter[n+i]= arr[i]
		} 
		for i := n - 1; i >= 1; i-- {
			treeIter[i] = math.Max(treeIter[2*i], treeIter[2*i+1])
		}
	}

	find = func(index, l, r, u, v int) float64 {
		if (u <= l && v >= r) {
			return tree[index]
		}
		mid := l + (r - l) / 2
		max := 0.0
		if u <= mid {
			max = math.Max(max, find( 2 * index + 1, l, mid, u, v))
		} 
		if v > mid {
			max = math.Max(max, find( 2 * index + 2, mid + 1, r, u, v))
		}
		return max
	}
	

	buildTree(0, 0, len(arr) - 1)
	buildTreeIter()
	fmt.Println(tree)
	fmt.Println(treeIter)
	fmt.Println(find(0, 0, len(arr) - 1, 2,9))
}

