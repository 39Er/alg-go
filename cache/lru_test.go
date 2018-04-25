package cache

import (
	"container/list"
	"fmt"
	"testing"
)

var cache *LRUCache

func TestLru(t *testing.T) {
	cache, _ = NewLRUCache(3)
	cache.Add("a", 1)
	cache.Add("b", 2)
	cache.Add("c", 3)
	print()

	cache.Get("a")
	cache.Add("d", 4)
	print()

	cache.Add("e", 5)
	cache.Add("f", 6)
	print()
}
func print() {
	printMap(cache.cache)
	printList(cache.ll)
	fmt.Println()
}
func printMap(cache map[Key]*list.Element) {
	var s string
	for k, v := range cache {
		s = fmt.Sprint(s, k, ":", v.Value, ", ")
	}
	fmt.Println("map:", s)
}

func printList(list *list.List) {
	ele := list.Front()
	strArr := []string{}
	for ele != nil {
		str := fmt.Sprint(ele.Value)
		strArr = append(strArr, str)
		ele = ele.Next()
	}
	fmt.Printf("list:%v\n", strArr)
}
