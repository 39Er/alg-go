package main

import "fmt"

//Byte Dance

type Node struct {
	value int
	next  *Node
}

func newNode(value int, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

func main() {
	tail := newNode(5, nil)
	node1 := newNode(4, tail)
	node2 := newNode(3, node1)
	node3 := newNode(2, node2)
	head := newNode(1, node3)
	printList(head)
	node := reverse2(head)
	printList(node)
}

func printList(node *Node) {
	fmt.Printf("%d\t", node.value)
	for node.next != nil {
		node = node.next
		fmt.Printf("%d\t", node.value)
	}
	fmt.Println()
}

//递归版本
func reverse1(node *Node) *Node {
	fmt.Println(node.value)
	if node == nil || node.next == nil {
		return node
	}
	result := reverse1(node.next)
	node.next.next = node
	node.next = nil
	return result
}

//迭代版本
func reverse2(node *Node) *Node {
	var next *Node
	var prev *Node
	for node != nil {
		next = node.next
		node.next = prev
		prev = node
		node = next
	}
	return prev
}
