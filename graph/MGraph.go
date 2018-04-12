package main

import "fmt"

const (
	MAXVEX   = 100
	INFINITY = 65535
)

//邻接矩阵
type MGraph struct {
	vexs                     [MAXVEX]string      //顶点表
	arc                      [MAXVEX][MAXVEX]int //矩阵表
	numVertexTypes, numEdges int                 //顶点数，边数
}

func CreateMGraph() *MGraph {
	G := new(MGraph)
	fmt.Println("输入顶点数和边数:")
	fmt.Scanf("%d,%d", &G.numVertexTypes, &G.numEdges)
	//初始化顶点表
	//fmt.Println("输入顶点名称：")
	for i := 0; i < G.numVertexTypes; i++ {
		//fmt.Scanf("%s", &G.vexs[i])
		G.vexs[i] = fmt.Sprintf("顶点%s", i)
	}
	//初始化矩阵
	for i := 0; i < G.numVertexTypes; i++ {
		for j := 0; j < G.numVertexTypes; j++ {
			G.arc[i][j] = INFINITY
		}
	}
	for k := 0; k < G.numEdges; k++ {
		var i, j, w int
		fmt.Println("输入边（vi,vj）上的下标i、下标j和权值w：")
		fmt.Scanf("%d,%d,%d", &i, &j, &w)
		G.arc[i][j] = w
		G.arc[j][i] = G.arc[i][j] //无向图，对称
	}
	return G
}
