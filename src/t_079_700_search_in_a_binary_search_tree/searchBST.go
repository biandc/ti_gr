package t_079_700_search_in_a_binary_search_tree

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

// root      = []int{4, 2, 7, 1, 3}
// val       = 2
func searchBST(root *TreeNode, val int) *TreeNode {
	var traversal func(node *TreeNode) *TreeNode
	traversal = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val > val {
			return traversal(node.Left)
		} else if node.Val < val {
			return traversal(node.Right)
		} else {
			return node
		}
	}
	return traversal(root)
}
