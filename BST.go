package main

import (
	"fmt"
)

type treeNode struct {
	Left  *treeNode
	Right *treeNode
	Value int
}

// insert a node into a existing tree
func (t *treeNode) Insert(data int) {
	if t == nil {
		t = &treeNode{Value: data}
		return
	}
	if data == t.Value {
		fmt.Println("Node already exists")
		return
	}
	if data < t.Value {
		if t.Left == nil {
			t.Left = &treeNode{Value: data}
			return
		}
		t.Left.Insert(data)
		return
	}
	if data > t.Value {
		if t.Right == nil {
			t.Right = &treeNode{Value: data}
			return
		}
		t.Right.Insert(data)
	}
}

// fetch the minimum element from the tree (fetch min)
func FetchMin(t *treeNode) int {
	if t.Left == nil {
		return t.Value
	}
	return FetchMin(t.Left)
}

// fetch the maximum element from the tree (fetch max)
func FetchMax(t *treeNode) int {
	if t.Right == nil {
		return t.Value
	}
	return FetchMax(t.Right)
}

// Node search
func (t *treeNode) nodeSearch(data int) *treeNode {
	if t == nil {
		fmt.Println("Tree is empty")
		return nil
	}
	if data == t.Value {
		return t
	}
	if data < t.Value {
		if t.Left == nil {
			fmt.Println("Element is not found")
			return nil
		}
		return t.Left.nodeSearch(data)
	}
	if t.Right == nil {
		fmt.Println("Element is not found")
		return nil
	}
	return t.Right.nodeSearch(data)
}

// Find parent node for a given node in BST
func (t *treeNode) parent(node *treeNode) *treeNode {
	if t == nil {
		fmt.Println("Tree does not exist")
		return nil
	}
	if t == node {
		fmt.Println("Node is root node, that has no parent")
		return nil
	}
	if node.Value < t.Value {
		if t.Left == node {
			return t
		}
		return t.Left.parent(node)
	}
	if t.Right == node {
		return t
	}
	return t.Right.parent(node)
}

// Traversal in a Binary Search Tree
// Preorder traversal
func (t *treeNode) PreOrderTraversal() {
	if t == nil {
		return
	}
	fmt.Println(t.Value)
	t.Left.PreOrderTraversal()
	t.Right.PreOrderTraversal()
}

//  Inorder traversal
func (t *treeNode) InOrderTraversal() {
	if t == nil {
		return
	}
	t.Left.InOrderTraversal()
	fmt.Println(t.Value)
	t.Right.InOrderTraversal()
}

//  PostOrder traversal
func (t *treeNode) PostOrderTraversal() {
	if t == nil {
		return
	}
	t.Left.PostOrderTraversal()
	t.Right.PostOrderTraversal()
	fmt.Println(t.Value)
}

// Breadth first search (level order from 0 to n)
func (t *treeNode) BreadthFirstTraversal() []int {
	if t == nil {
		return nil
	}
	var queue []*treeNode
	queue = append(queue, t)
	var res []int
	res = levelOrder(queue, res)
	fmt.Println(res)
	return res
}

func levelOrder(queue []*treeNode, res []int) []int {
	if len(queue) == 0 {
		return res
	}
	n := len(queue)
	for i := 0; i < n; i++ {
		res = append(res, queue[i].Value)
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}
	}
	return levelOrder(queue[n:], res)
}

// TODO Reverse level order Traversal

// print left view of BST
//Left view of a Binary Tree is set of nodes visible when tree is visited from left side.
// Supporting function: left
func (t *treeNode) printLeftView(root *treeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Value)
	left(root)
}
func left(t *treeNode) {
	if t == nil {
		return
	}
	if t.Left != nil {
		fmt.Println(t.Left.Value)
	}
	left(t.Left)
	left(t.Right)
}

// printLeftView1 - other version, supporting function: printLeft
func (t *treeNode) printLeftView1(root *treeNode) {
	if t == nil {
		return
	}
	var queue []*treeNode
	queue = append(queue, t)
	fmt.Println(t.Value)
	printLeft(queue)
}
func printLeft(queue []*treeNode) {
	if len(queue) == 0 {
		return
	}
	n := len(queue)
	for i := 0; i < n; i++ {
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
			fmt.Println(queue[i].Left.Value)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}
	}
	printLeft(queue[n:])
}

// read from right to left, print with height
func (t *treeNode) print(height int) {
	if t == nil {
		return
	}
	format := "--["
	t.Right.print(height + 1)
	fmt.Printf("%*s%d\n", 7*(height+1), format, t.Value)
	t.Left.print(height + 1)
}

// Find Inorder predecessor successor for a given key in BST
// Return -1 if there is not any
func (t *treeNode) FindPredecessor(key int) int {
	if t == nil {
		panic("root is nil")
	}
	node := t.nodeSearch(key)
	if node == nil {
		panic("key does not exist in the tree")
	}
	parent := t.parent(node)
	if node.Left == nil {
		if node.Value == t.Value {
			return -1
		}

		if node.Value < t.Value {
			if node.Value < parent.Value {
				return -1
			}
			return parent.Value
			//panic("No predecessor")
		}
		if node.Value > t.Value {
			if node.Value > parent.Value {
				return parent.Value
			}
			return t.parent(parent).Value
		}
	}
	return FetchMax(node.Left)
}

// Find Inorder successor for a given key in BST
func (t *treeNode) FindSuccessor(key int) int {
	if t == nil {
		panic("root is nil")
	}
	node := t.nodeSearch(key)
	if node == nil {
		panic("key does not exist in the tree")
	}
	parent := t.parent(node)
	if node.Right == nil {
		if node.Value == t.Value {
			return -1
		}

		if node.Value > t.Value {
			//panic("No successor")
			if node == parent.Right {
				return -1
			}
			return parent.Value
		}
		if node.Value < t.Value {
			if node == parent.Left {
				return parent.Value
			}
			return t.parent(parent).Value
		}
	}
	return FetchMin(node.Right)
}

/* ------------------------------------------------------------------------------
                                    DELETE
Delete a node from BST
Helper functions: PointToNil and transplant
---------------------------------------------------------------------------------*/
func (t *treeNode) DeleteNode(root *treeNode, data int) {
	if root == nil {
		return
	}
	node := root.nodeSearch(data)
	if node == nil {
		panic("key does not exist in the tree")
	}
	parent := t.parent(node)
	// delete leaf
	if node.Left == nil && node.Right == nil {
		PointToNil(node, parent)
		return
	}
	if node.Left == nil {
		t.transplant(node, node.Right)
		return
	}
	if node.Right == nil {
		t.transplant(node, node.Left)
		return
	}
	nodeSuccessor := root.FindSuccessor(data)
	node.Value = nodeSuccessor
	root.DeleteNode(node.Right, nodeSuccessor)
	return
}
func PointToNil(node, parent *treeNode) {
	if node == parent.Left {
		parent.Left = nil
		return
	}
	if node == parent.Right {
		parent.Right = nil
		return
	}
}
func (t *treeNode) transplant(n1, n2 *treeNode) {
	parent := t.parent(n1)
	if parent == nil {
		t.Value = n2.Value
		t.Right, t.Left = n2.Right, n2.Left
		return
	}
	if n1 == parent.Left {
		parent.Left = n2
		return
	}
	if n1 == parent.Right {
		parent.Right = n2
		return
	}
}

// get heights of all leaves and return them in an array, then find the max height using GetHeight function
func (t *treeNode) GetHeightLeaves(height int, leaves *[]int) []int {
	if t.Left == nil && t.Right == nil {
		*leaves = append(*leaves, height)
	} else {
		height++
		if t.Left != nil {
			t.Left.GetHeightLeaves(height, leaves)
		}
		if t.Right != nil {
			t.Right.GetHeightLeaves(height, leaves)
		}
	}
	return *leaves
}
func (t *treeNode) GetHeight(leaves []int) int {
	return max(leaves)
}

type levelValue map[int][]int

func (t *treeNode) GetLevel(level int, m *levelValue) levelValue {
	(*m)[level] = append((*m)[level], t.Value)
	if t.Left == nil && t.Right == nil {
		return *m
	}
	level = level + 1
	if t.Left == nil {
		return t.Right.GetLevel(level, m)
	}
	if t.Right == nil {
		return t.Left.GetLevel(level, m)
	}
	t.Left.GetLevel(level, m)
	t.Right.GetLevel(level, m)
	return *m
}

func main() {
	t := &treeNode{Value: 8}
	//t.print(0)
	values := []int{3, 10, 1, 6, 4, 7, 14, 13}
	for _,v := range values {
		t.Insert(v)
	}

	//t.PostOrderTraversal()
	//t.print(0)
	//fmt.Println(t.nodeSearch(12))
	//fmt.Println(t.FindPredecessor(8))
	//fmt.Println(t.FindSuccessor(5))
	//t.DeleteNode(t, 6)
	t.print(0)
	t.printLeftView(t)
	//leaf:= make([]int,0,10)
	//t.GetHeightLeaves(0,&leaf)
	//fmt.Println(leaf)
	//m := make(map[int][]int)
	//m := make(levelValue)
	//t.GetLevel(0, &m)
	//a := 1
	//fmt.Printf("Values at level %v: %v", a, m[a])
	//a := t.nodeSearch(2)
	//if a != nil {
	//	fmt.Println(a.Value)
	//}

}

func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}
