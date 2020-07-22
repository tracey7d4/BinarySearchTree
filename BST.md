# Binary Search Tree
A Binary Tree is a tree data structure, that consists of nodes. The topmost node in the tree is the root. 
Every node in a Binary Tree has maximum 2 children, which are referred to as left and right child. A leaf is 
a node that both its left and right subtrees are empty. A node with only one subtree is called half-leaf. 
A parent of a node is the node immediately above it. 

```go
type treeNode struct {
	Left  *treeNode
	Right *treeNode
	Value int
}
```
A common operations in a binary tree are searching, insertion, and deletion. To perform those operations in binary
search tree, we need to traverse all elements of the tree. Therefore, in the worst case, the time complexity of
those opereation is `O(n)`.

A Binary Search tree (BST) is a binary tree where values of the left node and right node are less than and greater
than value of its parent, correspondingly.

![binary search tree](image/binarySearchTree.png)

Operations in BST has time complexity is `O(h)`, where `h` is height of the BST. In the worst case where `h` is equal
number of elements in the tree, the worst case complexity is `O(n)`. 
A BST that is built from a sorted list of keys will produce worst case performance.

### Searching an element from BST
  Searching for an element from BST is performed by comparing the search element with the value of root node.
  * If both are matched, then return the root node
  * If search element is smaller, continue searching on the left subtree
  * If search element is larger, continue searching on the right subtree.
  * If there is no such element found, display "Element is not found" 

```go
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
```

### Insert an element to BST
Inserting an element to a BST with the same logic as searching for an element. It is noted that the element is
inserted to a leaf node only. 

```go
func (t *treeNode) Insert(data int)  {
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
```
### Fetch maximum value
Maximum value of a BST is the value of the most right node

```go
func FetchMax(t *treeNode) int {
	if t.Right == nil {
		return t.Value
	}
	return FetchMax(t.Right)
}
```
Return value: 14

### Fetch minimum value
Minimum value of a BST is the value of the most left node 

```go
func FetchMin(t *treeNode) int {
	if t.Left == nil {
		return t.Value
	}
	return FetchMin(t.Left)
}
```
Return: 1

### Traverse a BST
[A traversal](https://www.cs.cmu.edu/~adamchik/15-121/lectures/Trees/trees.html) 
is a process of visiting all the nodes in a tree.

We can categorise tree's traversal algorithms into 2 groups:  depth-first traversal and breadth-first traversal.
   * Depth-first traversal: includes PreOrder traversal, InOrder traversal, and PostOrder traversal.
   * Breadth-first traversal: the level order traversal. In this traversal, we visit the node by levels from
   top to bottom, and from left to right.

Let have a look at implementations of those traversal.

#### Preorder traversal
   0. Visit root node
   0. Visit all the nodes in the left subtree
   0. Visit all the nodes in the right subtree
   
   ```go
func (t *treeNode) PreorderTraversal() {
	if t == nil {
		return
	}
	fmt.Println(t.Value)
	t.Left.PreorderTraversal()
	t.Right.PreorderTraversal()
}
```
   Return: 8 3 1 6 4 7 10 14 13
   
#### Inorder traversal
   0. First, visit all the nodes in the left subtree
   0. Then the root node
   0. Visit all the nodes in the right subtree
 ```go
func (t *treeNode) InOrderTraversal() {
	if t == nil {
		return
	}
	t.Left.InOrderTraversal()
	fmt.Println(t.Value)
	t.Right.InOrderTraversal()
}
```  
   Return: 1 3 4 6 7 8 10 13 14
   
#### Postorder traversal
   0. Visit all the nodes in the left subtree
   0. Visit all the nodes in the right subtree
   0. Visit the root node
 
 ```go
func (t *treeNode) PostOrderTraversal() {
	if t == nil {
		return
	}
	t.Left.PostOrderTraversal()
	t.Right.PostOrderTraversal()
	fmt.Println(t.Value)
}
```  
   Return: 1 4 7 6 3 13 14 10 8 
   
#### Breadth first traversal
```go
func (t *treeNode) BreadthFirst() []int {
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
```
Return: 8 3 10 1 6 14 4 7 13

Now let's play a bit with our BST.

**Get Height**

Helper functions: `GetHeightLeaves` and `max`
```go
func (t *treeNode) GetHeight(leaves []int) int {
	return max(leaves)
}
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

func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}
```
One can also store an information about height in type node. In that case, height of a node is the maximum
height between left subtree and right subtree plus one.
```shell script
height = 1 + max(height_left, height_right)
```

**Print a tree**
```go
func (t *treeNode) print(height int) {
	if t == nil {
		return
	}
	format := "--["
	t.Right.print(height + 1)
	fmt.Printf("%*s%d\n", 7*(height+1), format, t.Value)
	t.Left.print(height + 1)
}
```
That is what will be displayed.

![Print tree](image/PrintTree.png)


**Get all values with the same level from BST**

   While the function `BreadthFirstTraversal` returns an array of tree's values by levels, the `GetLevel` function
  returns a map, where map's keys are tree's levels, and map's values are arrays of all the 
 nodes' values on each level.
 
```go
type levelValue map[int][]int
func (t *treeNode) GetLevel(level int, m *levelValue) levelValue {
	(*m)[level] = append((*m)[level],t.Value)
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
```
Such as
   0. Level 0: 8
   0. Level 1: 3 10
   0. Level 2: 1 6 14
   0. Level 3: 4 7 13
   
**Print left view of BST**

Left view of a Binary Tree is a set of nodes visible when tree is visited from left side.
```go
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
```
Return: 8 3 1 4 13

**Find Inorder predecessor and successor for a given key in BST**

In a given BST highest element on the left subtree is the Predecessor 
and lowest element on the right subtree is the Successor of the given node.
  0. Finding Predecessor
  ```go
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
```
  0. Finding Successor
  ```go
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
```
### Delete Node
   0. Node to be deleted is a leaf: Removing this node from the tree by setting its parent points to nil.
   0. Node to be deleted is half leaf: setting its parent points to its child
   0. Node to be deleted has two children: 
        * Find inorder successor of the node. 
        * Copy contents of the inorder successor to the node 
        * Delete the inorder successor. 
   It is noted that inorder predecessor can also be used.

```go
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
    // n1 is root
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
```
For example, deleting a leaf node - Node 1

![Print tree](image/DeleteNode1.png)

Deleting a half node - Node 14

![Print tree](image/DeleteNode14.png)

Deleting a full node - Node 6

![Print tree](image/DeleteNode6.png)

Deleting root node - Node 8

![Print tree](image/DeleteNode8.png)


### References

* [cs.cmu.edu](https://www.cs.cmu.edu/~adamchik/15-121/lectures/Trees/trees.html)