package main

import "fmt"

var visited []bool

func DFS(G *MGraph, i int) {
	visited[i] = true
	fmt.Println("经过顶点", G.vexs[i])
	for j := 0; j < G.numVertexTypes; j++ {
		//fmt.Println(i, "遍历到", j, G.arc[i][j] != INFINITY && !visited[j])
		if G.arc[i][j] != INFINITY && !visited[j] {
			DFS(G, j)
		}
	}
}

func DFSTraverse(G *MGraph) {
	visited = make([]bool, G.numVertexTypes)
	for i := 0; i < G.numVertexTypes; i++ {
		visited[i] = false
	}
	for i := 0; i < G.numVertexTypes; i++ {
		if !visited[i] {
			DFS(G, i)
		}
	}
}
