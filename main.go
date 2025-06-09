package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
	arr = mergeSort(arr)
	fmt.Println(arr)

	g := new_graph(3)
	dfs_graph(g)
	fmt.Println()
}
