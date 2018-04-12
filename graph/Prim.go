package main

import "fmt"

/*Prim算法生成最小生成树*/
/*
	1、默认将0顶点设为起始，将其他顶点到v0的距离保存在lowcost中，adjvex中关联的顶点都设为是0
	2、找出lowcost中的最小值（已加入到最小生成树中的顶点不需要再进行比较），将lowcost中的相应值改为0，表示此顶点k已加入到最小生成树种
	3、将其他顶点到顶点k的距离与现在相应lowcost中的值进行比较替换，并将adjvex中的关联的顶点设为相应的最小边权值的顶点
	4、再次循环遍历
*/

func Prim(G *MGraph) {
	var (
		adjvex  [MAXVEX]int //保存相关顶点下标
		lowcost [MAXVEX]int //保存相关顶点间边的权值
		j, k    int
	)
	lowcost[0] = 0
	adjvex[0] = 0
	for i := 1; i < G.numVertexTypes; i++ {
		lowcost[i] = G.arc[0][i]
		adjvex[i] = 0
	}
	for i := 1; i < G.numVertexTypes; i++ {
		min := INFINITY
		j, k = 1, 0
		for j < G.numVertexTypes {
			if lowcost[j] != 0 && min > lowcost[j] { //已加入到最小生成树中的顶点不需要再进行比较
				min = lowcost[j]
				k = j
			}
			j++
		}
		fmt.Println(adjvex[k], k)
		lowcost[k] = 0
		for j = 1; j < G.numVertexTypes; j++ {
			if lowcost[j] != 0 && lowcost[j] > G.arc[k][j] { //已加入到最小生成树中的顶点不需要再进行比较
				lowcost[j] = G.arc[k][j]
				adjvex[j] = k
			}
		}
	}
}
