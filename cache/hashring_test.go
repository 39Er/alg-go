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
		nodes[nodeKey] = i * 10
	}
	hashring.AddNodes(nodes)
	result := make(map[string]int)
	addResult := func(key string) {
		if value, ok := result[key]; !ok {
			result[key] = 1
		} else {
			result[key] = value + 1
		}
	}

	for i := 0; i < 1000; i++ {
		addResult(hashring.GetNode(string(i)))
	}

	fmt.Println(result)

	fmt.Println("----------------------")

	result = make(map[string]int)
	hashring.Delete("node1")
	for i := 0; i < 1000; i++ {
		addResult(hashring.GetNode(string(i)))
	}
	fmt.Println(result)
}
