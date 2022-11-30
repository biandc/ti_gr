package t_061_226_invert_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		invert(node.Left)
		invert(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	invert(root)
	return root
}
