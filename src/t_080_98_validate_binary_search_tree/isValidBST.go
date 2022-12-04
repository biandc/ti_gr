package t_080_98_validate_binary_search_tree

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

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var traversal func(node *TreeNode) bool
	traversal = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := traversal(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		} else {
			pre = node
		}
		right := traversal(node.Right)
		return left && right
	}
	return traversal(root)
}
