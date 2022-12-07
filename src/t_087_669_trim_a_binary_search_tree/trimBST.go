package t_087_669_trim_a_binary_search_tree

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

// 递归
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var traversal func(cur *TreeNode) *TreeNode
	traversal = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}
		if cur.Val < low {
			return traversal(cur.Right)
		}
		if cur.Val > high {
			return traversal(cur.Left)
		}
		cur.Left = traversal(cur.Left)
		cur.Right = traversal(cur.Right)
		return cur
	}
	return traversal(root)
}

