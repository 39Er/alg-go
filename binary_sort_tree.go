package alg_go

import "fmt"

type BiTree struct {
	data           int
	lchild, rchild *BiTree
}

func CreateBST(arr []int) *BiTree {
	var tree *BiTree
	for _, i := range arr {
		tree = InsertBST(tree, i)
	}
	return tree
}

func InsertBST(tree *BiTree, key int) *BiTree {
	if tree == nil {
		tree = &BiTree{key, nil, nil}
		return tree
	}
	if SearchBST(tree, key) {
		return tree
	}

	if tree.data > key {
		tree.lchild = InsertBST(tree.lchild, key)
	} else {
		tree.rchild = InsertBST(tree.rchild, key)
		//tree.rchild = &BiTree{key, nil, nil}
	}
	return tree
}

func SearchBST(tree *BiTree, key int) bool {
	if tree == nil {
		return false
	}
	if tree.data == key {
		return true
	} else if tree.data > key {
		return SearchBST(tree.lchild, key)
	} else {
		return SearchBST(tree.rchild, key)
	}
}

func ListBST(tree *BiTree) {
	if tree == nil {
		return
	}
	ListBST(tree.lchild)
	fmt.Printf("%d ", tree.data)
	ListBST(tree.rchild)
}

func DelBST(tree *BiTree, key int) bool {
	if tree == nil {
		return false
	}
	if tree.data == key {
		deleteNode(tree)
		return true
	}
	if tree.data < key {
		return DelBST(tree.rchild, key)
	} else {
		return DelBST(tree.lchild, key)
	}
}

func deleteNode(tree *BiTree) {
	if tree.lchild == nil {
		*tree = *tree.rchild
		return
	}
	if tree.rchild == nil {
		*tree = *tree.lchild
		return
	}
	previous := tree
	next := tree.lchild
	for next.rchild != nil {
		previous = next
		next = next.rchild
	}
	tree.data = next.data
	if previous.lchild == next { //若节点tree的左子树没有右子树,此时previous=tree
		previous.lchild = next.lchild
	} else {
		previous.rchild = next.rchild
	}
}
