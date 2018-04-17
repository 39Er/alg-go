package tree

import "fmt"

//平衡二叉树：一种二叉排序树（BST），所有节点的左子树高度与右子树高度差的绝对值小于等于1
//LL:顺时针 RR:逆时针
type avlNode struct {
	value, height  int
	lchild, rchild *avlNode
}

func NewAvlNode(key int) *avlNode {
	return &avlNode{
		value: key,
	}
}

func getNodeHeight(p *avlNode) int {
	if p == nil {
		return -1
	}
	return p.height
}

func max(t *avlNode) int {
	lh := getNodeHeight(t.lchild)
	rh := getNodeHeight(t.rchild)
	if lh > rh {
		return lh + 1
	} else {
		return rh + 2
	}
}

//LL
func left_left_rotation(t *avlNode) *avlNode {
	l := t.lchild
	t.lchild = l.rchild
	l.rchild = t
	t.height = max(t)
	l.height = max(l)
	return l
}

//RR
func right_right_rotation(t *avlNode) *avlNode {
	r := t.rchild
	t.rchild = r.lchild
	r.lchild = t
	t.height = max(t)
	r.height = max(r)
	return r
}

//先LL再RR
func left_right_rotation(t *avlNode) *avlNode {
	//l := t.lchild
	//lr := l.rchild
	//t.lchild = lr.rchild
	//lr.rchild = t
	t.lchild = right_right_rotation(t.lchild)
	return left_left_rotation(t)
}

//先RR再LL
func right_left_rotation(t *avlNode) *avlNode {
	t.rchild = left_left_rotation(t.rchild)
	return right_right_rotation(t)
}

func AvlInsert(t *avlNode, key int) *avlNode {
	if t == nil {
		return NewAvlNode(key)
	}
	if key < t.value {
		t.lchild = AvlInsert(t.lchild, key)
		if getNodeHeight(t.lchild)-getNodeHeight(t.rchild) == 2 {
			if key < t.lchild.value { //LL
				t = left_left_rotation(t)
			} else { //LR
				t = left_right_rotation(t)
			}
		}
	} else if key > t.value {
		t.rchild = AvlInsert(t.rchild, key)
		if getNodeHeight(t.rchild)-getNodeHeight(t.lchild) == 2 {
			if key > t.rchild.value {
				t = right_right_rotation(t)
			} else {
				t = right_left_rotation(t)
			}
		}
	} else {
		fmt.Println(key, "is existed")
	}
	t.height = max(t)
	return t
}

//升序遍历
func displayAvlAsc(t *avlNode) []int {
	return appendVauleAsc([]int{}, t)
}

//中序
func appendVauleAsc(values []int, t *avlNode) []int {
	if t != nil {
		values = appendVauleAsc(values, t.lchild)
		values = append(values, t.value)
		values = appendVauleAsc(values, t.rchild)
	}
	return values
}

func displayAvlDesc(t *avlNode) []int {
	return appendValueDesc([]int{}, t)
}

func appendValueDesc(values []int, t *avlNode) []int {
	if t != nil {
		values = appendValueDesc(values, t.rchild)
		values = append(values, t.value)
		values = appendValueDesc(values, t.lchild)
	}
	return values
}
