package main

import(
	"fmt"
)

/**
You have planned some train traveling one year in advance. The days of the year in which 
you will travel are given as an integer array days, Each day is an integer from 1 to 365.
Train tickets are sold in three different ways.
- a 1-day pass is sold for cost[0] dollars.
- a 7-day pass is sold for cost[1] dollars
- a 30-day pass is sold for cost[2] dollars

The passes allow that many days of consecutive travel.
- For example, if we get a 7-day on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.
Return the minimum number of dollars you need to travel every day in the given list of days.

*/
var graph map[int][]Point

type Point struct {
	Destination int
	Distance int
}

func main() {
	days := []int {1, 2, 3, 4, 5, 6,7, 10,11,12,13, 20,31}
	costs := []int{2, 7, 15}
	n := len(days)
	graph = buildGraph(days)
	var dfs func(int, int) int 

	dfs = func(pos, cost int) int {
		if pos >= n {
			return cost
		}
		op1 := dfs(travel(pos, 1), cost + costs[0])
		op2 := dfs(travel(pos, 7), cost + costs[1])
		op3 := dfs(travel(pos, 30), cost + costs[2])
		return min(op1, op2, op3)
	}
	fmt.Println(dfs(0,0))
}

func buildGraph(days []int) map[int][]Point {
	graph = make(map[int][]Point)
	for i := 0; i < len(days); i++ {
		for j := i; j < len(days); j++ {
			graph[i] = append(graph[i], Point {Destination: j, Distance: days[j]-days[i] + 1})
		}
	}
	return graph
}

func travel(from, distance int) int {
	for _, v := range graph[from] {
		if v.Distance <= distance {
			from = v.Destination + 1
		} else {
			break
		}
	}
	return from
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}