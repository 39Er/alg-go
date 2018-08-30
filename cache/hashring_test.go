package cache

import (
	"fmt"
	"testing"
)

func TestHashRing_AddNode(t *testing.T) {
	hashring := NewHashRing()
	hashring.AddNode("node0", 1)
	nodes := make(map[string]int)
	for i := 1; i < 5; i++ {
		nodeKey := fmt.Sprintf("node%d", i)
		nodes[nodeKey] = i
	}
	hashring.AddNodes(nodes)
	for i := 0; i < 10; i++ {
		fmt.Println(hashring.GetNode(string(i)))
	}

	fmt.Println("----------------------")
	hashring.Delete("node1")
	for i := 0; i < 10; i++ {
		fmt.Println(hashring.GetNode(string(i)))
	}
}
