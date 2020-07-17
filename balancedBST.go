package main

import "fmt"

type Node struct {
	Value int   // value of a node
	Left  *Node // left node
	Right *Node // right node
	Bal   int   // Bal = h_left - h_right
}

/*--------------------------------------------------------------------------------
		                   INSERTION
Insert a given data to an existing balanced subtree at node n,
re-balance the subtree by calling Rebalance function if needed
--------------------------------------------------------------------------------*/
func (n *Node) Insert(data int) {
	// Node already exists, do no change
	if data == n.Value {
		fmt.Println("Node already exists")
		return
	}
	// If data is less than value of node c, traversal to the left, otherwise, traversal to the right
	if data < n.Value {
		// left node does not exist, create a new left node
		if n.Left == nil {
			n.Left = &Node{Value: data}
			//update Balance factor of node n
			n.Bal++
			return
		}

		a := n.Left.Bal
		n.Left.Insert(data)
		// check whether the tree's height has changed after inserting
		if n.Left.Bal != 0 && n.Left.Bal != a {
			if n.Left.Bal < -1 || n.Left.Bal > 1 {
				n.Rebalance(n.Left)
			} else {
				n.Bal++
			}
		}
		return
	}

	if n.Right == nil {
		n.Right = &Node{Value: data}
		n.Bal--
		return
	}

	b := n.Right.Bal
	n.Right.Insert(data)
	if n.Right.Bal != 0 && n.Right.Bal != b {
		if n.Right.Bal < -1 || n.Right.Bal > 1 {
			n.Rebalance(n.Right)
		} else {
			n.Bal--
		}
	}
	return
}

/*--------------------------------------------------------------------------------
		                   RE-BALANCING
Perform suitable rotation operation to make the subtree balanced.
c is considered node
n is parent node of node c
Helper functions: LLRotation, RRRotation, LRRotation, RLRotation
--------------------------------------------------------------------------------*/

func (n *Node) Rebalance(c *Node) {
	fmt.Printf("Re-Balance @ node %v \n", c.Value)
	// Left subtree is too high, and left child has a left child.
	if c.Bal == 2 && c.Left.Bal == 1 {
		n.RRRotation(c)
		return
	}
	// Right subtree is too high, and right child has a right child.
	if c.Bal == -2 && c.Right.Bal == -1 {
		n.LLRotation(c)
		return
	}
	// Left subtree is too high, and left child has a right child.
	if c.Bal == 2 && c.Left.Bal == -1 {
		n.LRRotation(c)
		return
	}
	// Right subtree is too high, and right child has a left child.
	if c.Bal == -2 && c.Right.Bal == 1 {
		n.RLRotation(c)
		return
	}
}

func (n *Node) LLRotation(c *Node) {
	r := c.Right     // create a temporary node r to hold node c.Right
	c.Right = r.Left //  make  left children of r become right children of c
	r.Left = c       // make c as r left child
	// update balance factor of node c and r
	if c.Left != nil && c.Right == nil {
		c.Bal = 1
	} else {
		c.Bal = 0
	}
	r.Bal = 0
	// let parent node n point to new node r
	if c == n.Left {
		n.Left = r
	} else {
		n.Right = r
	}
}

func (n *Node) RRRotation(c *Node) {
	l := c.Left      // create a temporary node l to hold node c.Left
	c.Left = l.Right //  make left children of l become right children of c
	l.Right = c      // make c as l's right child

	if c.Right != nil && c.Left == nil {
		c.Bal = -1
	} else {
		c.Bal = 0
	}
	l.Bal = 0
	// let parent node n point to new node l
	if c == n.Left {
		n.Left = l
	} else {
		n.Right = l
	}
}

func (n *Node) LRRotation(c *Node) {
	// Left Rotation at node c.Left
	c.LLRotation(c.Left)
	// Right Rotation at node c
	n.RRRotation(c)
}

func (n *Node) RLRotation(c *Node) {
	// Right Rotation at node c.Right
	c.RRRotation(c.Right)
	// Left Rotation at node c
	n.LLRotation(c)
}

/*--------------------------------------------------------------------------------
		                   TREE TYPE
A tree can either be empty or has a node. Each node in a tree has left subtree, and right subtree.
If a tree is empty, inserting an element to the tree is creating a root node.
Performing insertion to a tree has to take care of re-balancing the root node if necessary
--------------------------------------------------------------------------------*/

type AVLTree struct {
	Root *Node
}

// Insertion
func (t *AVLTree) Insert(data int) {
	if t.Root == nil {
		t.Root = &Node{Value: data}
		return
	}
	t.Root.Insert(data)
	// Check whether the root node gets out of balance
	if t.Root.Bal < -1 || t.Root.Bal > 1 {
		t.rebalance()
	}
}

// The root node has no parent, to be unable to use node's Rebalance method above
// a fake parent node is created. This node stores value -1, and root node is its either
// left child or right child
func (t *AVLTree) rebalance() {
	rootParent := &Node{Left: t.Root, Value: -1}
	rootParent.Rebalance(t.Root)
	t.Root = rootParent.Left
}

//Print the tree with displaying nodes' heights
func (t *AVLTree) print(n *Node, height int) {
	if n == nil {
		return
	}
	format := "--["
	t.print(n.Right, height+1)
	fmt.Printf("%*s%d(%v)\n", 7*(height+1), format, n.Value, n.Bal)
	t.print(n.Left, height+1)
}

func main() {
	//values := []int{2, 4, 3, 5, 6, 1, 7, 8}
	//values := []int{7, 8, 9, 10, 11, 12, 13, 14, 1, 2, 3, 4, 5, 6}
	//values := []int{1,4,3}
	values := []int{11, 2, 13, 4, 3, 2, 5, 6, 1, 7, 8, 10, 9}
	t := &AVLTree{}
	for i := 0; i < len(values); i++ {
		fmt.Printf("Insert %v------------------------------------------\n", values[i])
		t.Insert(values[i])
		t.print(t.Root, 0)
		fmt.Println()
	}
}
