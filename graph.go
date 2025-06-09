package main

import "fmt"

type vertex struct {
	value   int
	edges   []vertex
	visited bool
}

type graph struct {
	vertecies_amount int
	vertecies        []vertex
}

func create_vertex(value int) vertex {
	v := vertex{value: value, visited: false}
	return v
}

func new_graph(vertecies_amount int) graph {
	vertecies := make([]vertex, 0, vertecies_amount)
	for i := 0; i < vertecies_amount; i++ {
		vertecies = append(vertecies, create_vertex(i))
	}

	return graph{
		vertecies_amount: vertecies_amount,
		vertecies:        vertecies,
	}
}

func dfs(v *vertex) {
	fmt.Printf("%d ", v.value)
	v.visited = true

	for _, e := range v.edges {
		if !e.visited {
			dfs(&e)
		}
	}
}

func dfs_graph(g graph) {
	for _, v := range g.vertecies {
		if !v.visited {
			dfs(&v)
		}
	}
}
