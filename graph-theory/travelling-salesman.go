package main

import(
	"fmt"
)

const MAX_INT = 1000000

func main() {
	graph := [][]int { { 0, 10, 15, 20 },{ 10, 0, 35, 25 },{ 15, 35, 0, 30 },{ 20, 25, 30, 0 } };
	start := 0;
	tsp(graph, start)
}

func tsp( graph [][]int, start int) {
	n := len(graph)

	//bitPresent := 1 << (n + 1)
	memo := setup(graph, start, n)
	fmt.Printf("%v \n", memo)

	memo = solve(graph, memo, start, n)
	minTourCost := findMinCost(graph, memo, start, n) 
	findOptimationTour(graph, memo, start, n) 
	fmt.Println("ssss: ", minTourCost)
}

func setup(graph [][]int, start int, n int) [5][32]int {
	memo := [5][32]int{};
	for i := 0; i < n; i ++ {
		if i == start {
			continue
		}
		// 1 << start | 1 << i => Mean: it show that moving from (start -> i) = distance.  
		memo[i][1 << start | 1 << i] = graph[start][i]
	}
	return memo
}

func solve(graph [][]int, memo [5][32]int, start int, n int ) [5][32]int {
	for r := 3; r <= n ; r ++ {
		// combinations(3, 4) = {0111, 1011, 1101, 1110}
		for _, subset := range combination(r, n) {
			// Subset mean is the route, ex: 0111: 0 -> 1 -> 2 | 0 -> 2 -> 1, 1011: 0 -> 1 -> 3...
			// the goal of below code is to find the minimum distance of Subset.

			if notIn(start, subset) {
				continue
			}
			for next := 0; next < n; next ++ {
				if next == start || notIn(next, subset) {
					continue
				}
				// State is like a child of subset
				// Ex: subset: 1011, state: 1001 => route: 0 -> 3 -> 1, state: 0011 => route: 0 -> 1 -> 3
				// subset: 1111, state: 1101 => route: 0 -> 2 -> 3 -> 1 | 0 -> 3 -> 2 -> 1
				// e like a destination after start.
				// next like a destination after e.
				state := subset ^ (1 << next)
				minDist := MAX_INT
				for e := 0; e < n; e ++ {
					fmt.Printf("subset: %b, state: %b, next: %d, e: %d  \n", subset,state, next, e)
					if e == start || e == next || notIn(e, subset) {
						continue
					}
					newDistance := memo[e][state] + graph[e][next]
					fmt.Printf("newDistance: (%d) := memo[e][state](%d) + graph[e][next] (%d) \n", newDistance,memo[e][state],  graph[e][next] )
					if newDistance < minDist {
						minDist = newDistance
					}
				}
				fmt.Printf("------memo[next(%d)][subset(%b)](%d) = minDist(%d) \n",next, subset, memo[next][subset], minDist )
				memo[next][subset] = minDist
			}
		}
	}
	return memo
}

var subsets []int

func combination(r, n int) []int{
	subsets = []int{}
	combinations(0, 0, r, n)
	//fmt.Printf("%b \n", subsets)
	return subsets
}

func combinations(set, at, r, n int) {
	if r == 0 {
		subsets = append(subsets, set)
	} else {
		for i := at; i < n; i ++ {
			set = set | (1 << i)
			combinations(set, i + 1, r - 1, n)
			set = set & (^(1 << i))
		} 
	}
}

func notIn(s int, subset int) bool {
	return ((1 << s) & subset) == 0
}

func findMinCost(graph [][]int, memo [5][32]int, start, n int) int{
	end_state := (1 << n) - 1
	mintourCost := MAX_INT

	for e := 0; e < n; e ++ {
		if e == start {
			continue
		}
		fmt.Printf("%b \n", end_state)
		tourCost := memo[e][end_state] + graph[e][start]
		if tourCost < mintourCost {
			mintourCost = tourCost
		}
	}
	return mintourCost
}

func findOptimationTour(graph [][]int, memo [5][32]int, start, n int) {
	lastIndex := start
	state := (1 << n) - 1 // End state
	tour := make([]int, 5)
	for i := n - 1; i >= 1; i-- {
		index := -1
		for j := 0; j < n; j++ {
			if j == start || notIn(j, state) {
				continue
			}
			if index == -1 {
				index = j
			}
			prevDist := memo[index][state] + graph[index][lastIndex]
			newDist := memo[j][state] + graph[j][lastIndex]
			
			if newDist < prevDist {
				index = j
			}
		}

		tour[i] = index
		state = state ^ (1 << index)
		lastIndex = index
	}
	tour[0] = start
	tour[n] = start
	fmt.Println(tour)
}