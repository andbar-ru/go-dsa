package dsa

import "fmt"

/**************************************************
 * Tree Node
 **************************************************/
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	height int // for AVL tree
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("%d", n.Val)
}

func (n *TreeNode) inorderTraversal() *List {
	if n == nil {
		return nil
	}
	dummy := new(ListNode)
	curNode := dummy

	// First, visit all the nodes in the left subtree
	leftTraversal := n.Left.inorderTraversal()
	if leftTraversal != nil {
		curNode.Next = leftTraversal.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
	}

	// Then the root node
	curNode.Next = &ListNode{Val: n.Val}
	curNode = curNode.Next

	// Visit all the nodes in the tight subtree
	rightTraversal := n.Right.inorderTraversal()
	if rightTraversal != nil {
		curNode.Next = rightTraversal.Head
	}

	return &List{Head: dummy.Next}
}

func (n *TreeNode) preorderTraversal() *List {
	if n == nil {
		return nil
	}
	dummy := new(ListNode)
	curNode := dummy

	// Visit n node
	curNode.Next = &ListNode{Val: n.Val}
	curNode = curNode.Next

	// Visit all the nodes in the left subtree
	leftTraversal := n.Left.preorderTraversal()
	if leftTraversal != nil {
		curNode.Next = leftTraversal.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
	}

	// Visit all the nodes in the tight subtree
	rightTraversal := n.Right.preorderTraversal()
	if rightTraversal != nil {
		curNode.Next = rightTraversal.Head
	}

	return &List{Head: dummy.Next}
}

func (n *TreeNode) postorderTraversal() *List {
	if n == nil {
		return nil
	}
	dummy := new(ListNode)
	curNode := dummy

	// Visit all the nodes in the left subtree
	leftTraversal := n.Left.postorderTraversal()
	if leftTraversal != nil {
		curNode.Next = leftTraversal.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
	}

	// Visit all the nodes in the tight subtree
	rightTraversal := n.Right.postorderTraversal()
	if rightTraversal != nil {
		curNode.Next = rightTraversal.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
	}

	// Visit n node
	curNode.Next = &ListNode{Val: n.Val}

	return &List{Head: dummy.Next}
}

// Insert on the left of the node
func (n *TreeNode) InsertLeft(val int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: val}
	}
	n.Left = &TreeNode{Val: val}
	return n.Left
}

// Insert on the right of the node
func (n *TreeNode) InsertRight(val int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: val}
	}
	n.Right = &TreeNode{Val: val}
	return n.Right
}

func (n *TreeNode) isFull() bool {
	// Checking n emptiness
	if n == nil {
		return true
	}

	// Checking the presence of children
	if n.Left == nil && n.Right == nil {
		return true
	}

	if n.Left != nil && n.Right != nil {
		return n.Left.isFull() && n.Right.isFull()
	}

	return false
}

func (n *TreeNode) leftmostDepth() int {
	d := 0
	for n != nil {
		d++
		n = n.Left
	}
	return d
}

func (n *TreeNode) isPerfect(d, level int) bool {
	// Check if the tree is empty
	if n == nil {
		return true
	}

	// Check the presence of children
	if n.Left == nil && n.Right == nil {
		return d == level+1
	}
	if n.Left == nil || n.Right == nil {
		return false
	}

	return n.Left.isPerfect(d, level+1) && n.Right.isPerfect(d, level+1)
}

func (n *TreeNode) countNodes() int {
	if n == nil {
		return 0
	}
	return 1 + n.Left.countNodes() + n.Right.countNodes()
}

func (n *TreeNode) isComplete(index, numNodes int) bool {
	if n == nil {
		return true
	}
	if index >= numNodes {
		return false
	}
	return n.Left.isComplete(2*index+1, numNodes) && n.Right.isComplete(2*index+2, numNodes)
}

func (n *TreeNode) isBalanced(height *int) bool {
	if n == nil {
		*height = 0
		return true
	}

	var leftHeight, rightHeight int
	var l, r bool
	l = n.Left.isBalanced(&leftHeight)
	r = n.Right.isBalanced(&rightHeight)

	if leftHeight > rightHeight {
		*height = leftHeight + 1
	} else {
		*height = rightHeight + 1
	}

	diff := leftHeight - rightHeight
	if abs(diff) >= 2 {
		return false
	} else {
		return l && r
	}
}

func (n *TreeNode) insertBST(val int) *TreeNode {
	// Return a new node if the tree is empty
	if n == nil {
		return &TreeNode{Val: val}
	}

	// Traverse to the right place and insert the node
	if val < n.Val {
		n.Left = n.Left.insertBST(val)
	} else {
		n.Right = n.Right.insertBST(val)
	}

	return n
}

func (n *TreeNode) leftmostLeaf() *TreeNode {
	curNode := n
	for curNode != nil && curNode.Left != nil {
		curNode = curNode.Left
	}
	return curNode
}

func (n *TreeNode) deleteBST(val int) *TreeNode {
	// Return if the tree is empty
	if n == nil {
		return n
	}

	// Find the node to be deleted
	if val < n.Val {
		n.Left = n.Left.deleteBST(val)
	} else if val > n.Val {
		n.Right = n.Right.deleteBST(val)
	} else {
		// If the node is with only one child or no child
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// If the node has two children
		inorderSuccessor := n.Right.leftmostLeaf()

		// Place the inorder successor in position of the node to be deleted
		n.Val = inorderSuccessor.Val

		// Delete the inorder successor
		n.Right.deleteBST(inorderSuccessor.Val)
	}

	return n
}

func (n *TreeNode) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Right rotate
func (y *TreeNode) rightRotateAVL() *TreeNode {
	x := y.Left
	t2 := x.Right
	x.Right = y
	y.Left = t2
	y.height = max(y.Left.Height(), y.Right.Height()) + 1
	x.height = max(x.Left.Height(), x.Right.Height()) + 1
	return x
}

// Left rotate
func (x *TreeNode) leftRotateAVL() *TreeNode {
	y := x.Right
	t2 := y.Left
	y.Left = x
	x.Right = t2
	x.height = max(x.Left.Height(), x.Right.Height()) + 1
	y.height = max(y.Left.Height(), y.Right.Height()) + 1
	return y
}

// Get the balance factor
func (n *TreeNode) balanceAVL() int {
	if n == nil {
		return 0
	}
	return n.Left.Height() - n.Right.Height()
}

// Insert node in an AVL tree
func (n *TreeNode) insertAVL(val int) *TreeNode {
	// Find the correct position to insert the node and insert it
	if n == nil {
		return &TreeNode{Val: val}
	}

	if val < n.Val {
		n.Left = n.Left.insertAVL(val)
	} else if val > n.Val {
		n.Right = n.Right.insertAVL(val)
	} else {
		return n
	}

	// Update the balance factor of each node and balance the n
	n.height = 1 + max(n.Left.Height(), n.Right.Height())
	balance := n.balanceAVL()
	if balance > 1 && val < n.Left.Val {
		return n.rightRotateAVL()
	}
	if balance < -1 && val > n.Right.Val {
		return n.leftRotateAVL()
	}
	if balance > 1 && val > n.Left.Val {
		n.Left = n.Left.leftRotateAVL()
		return n.rightRotateAVL()
	}
	if balance < -1 && val < n.Right.Val {
		n.Right = n.Right.rightRotateAVL()
		return n.leftRotateAVL()
	}

	return n
}

// Delete node in an AVL tree
func (n *TreeNode) deleteAVL(val int) *TreeNode {
	if n == nil {
		return n
	}

	if val < n.Val {
		n.Left = n.Left.deleteAVL(val)
	} else if val > n.Val {
		n.Right = n.Right.deleteAVL(val)
	} else {
		if n.Left == nil || n.Right == nil {
			var temp *TreeNode
			if n.Left != nil {
				temp = n.Left
			} else {
				temp = n.Right
			}

			if temp == nil {
				temp = n
				n = nil
			} else {
				n = temp
			}
		} else {
			temp := n.Right.leftmostLeaf()
			n.Val = temp.Val
			n.Right = n.Right.deleteAVL(temp.Val)
		}
	}

	if n == nil {
		return n
	}

	// Update the balance factor of each node and balance the tree
	n.height = 1 + max(n.Left.Height(), n.Right.Height())
	balance := n.balanceAVL()
	if balance > 1 && n.Left.balanceAVL() >= 0 {
		return n.rightRotateAVL()
	}
	if balance > 1 && n.Left.balanceAVL() < 0 {
		n.Left = n.Left.leftRotateAVL()
		return n.rightRotateAVL()
	}
	if balance < -1 && n.Right.balanceAVL() <= 0 {
		return n.leftRotateAVL()
	}
	if balance < -1 && n.Right.balanceAVL() > 0 {
		n.Right = n.Right.rightRotateAVL()
		return n.leftRotateAVL()
	}

	return n
}

/**************************************************
 * Binary Tree
 **************************************************/
type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{
		Root: &TreeNode{
			Val: val,
		},
	}
}

func (tree *BinaryTree) InorderTraversal() *List {
	if tree == nil {
		return nil
	}
	return tree.Root.inorderTraversal()
}

func (tree *BinaryTree) PreorderTraversal() *List {
	if tree == nil {
		return nil
	}
	return tree.Root.preorderTraversal()
}

func (tree *BinaryTree) PostorderTraversal() *List {
	if tree == nil {
		return nil
	}
	return tree.Root.postorderTraversal()
}

func (tree *BinaryTree) IsFull() bool {
	if tree == nil {
		return true
	}
	return tree.Root.isFull()
}

func (tree *BinaryTree) IsPerfect() bool {
	if tree == nil {
		return true
	}
	root := tree.Root
	d := root.leftmostDepth()
	return root.isPerfect(d, 0)
}

func (tree *BinaryTree) IsComplete() bool {
	if tree == nil {
		return true
	}
	root := tree.Root
	numNodes := root.countNodes()
	return root.isComplete(0, numNodes)
}

// Check if the tree is balanced
func (tree *BinaryTree) IsBalanced() bool {
	if tree == nil {
		return true
	}
	height := 0
	return tree.Root.isBalanced(&height)
}

/**************************************************
 * Binary Search Tree
 **************************************************/
type BinarySearchTree struct {
	*BinaryTree
}

func NewBinarySearchTree(val int) *BinarySearchTree {
	tree := NewBinaryTree(val)
	return &BinarySearchTree{tree}
}

// Insert a node into binary search tree (BST)
func (tree *BinarySearchTree) Insert(val int) *TreeNode {
	if tree == nil {
		return &TreeNode{Val: val}
	}
	return tree.Root.insertBST(val)
}

// Delete a node in the binary search tree (BST)
func (tree *BinarySearchTree) Delete(val int) *TreeNode {
	if tree == nil {
		return nil
	}
	return tree.Root.deleteBST(val)
}

/**************************************************
 * AVL Tree
 **************************************************/
type AVLTree struct {
	*BinarySearchTree
}

func NewAVLTree(val int) *AVLTree {
	tree := NewBinarySearchTree(val)
	return &AVLTree{tree}
}

// Insert node
func (tree *AVLTree) Insert(val int) *TreeNode {
	if tree == nil {
		return &TreeNode{Val: val}
	}
	root := tree.Root.insertAVL(val)
	tree.Root = root
	return root
}

// Delete node
func (tree *AVLTree) Delete(val int) *TreeNode {
	if tree == nil {
		return nil
	}
	root := tree.Root.deleteAVL(val)
	tree.Root = root
	return root
}
