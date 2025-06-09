package main

import "fmt"

type vertex struct {
	value   int
	edges   []vertex
	visited bool
}

type graph struct {
	verteciesAmount int
	vertecies        []vertex
}

func createVertex(value int) vertex {
	v := vertex{value: value, visited: false}
	return v
}

func newGraph(verteciesAmount int) graph {
	vertecies := make([]vertex, 0, verteciesAmount)
	for i := 0; i < verteciesAmount; i++ {
		vertecies = append(vertecies, createVertex(i))
	}

	return graph{
		verteciesAmount: verteciesAmount,
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

func dfsGraph(g graph) {
	for _, v := range g.vertecies {
		if !v.visited {
			dfs(&v)
		}
	}
}
