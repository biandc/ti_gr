package t_059_104_maximum_depth_of_binary_tree

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

func maxDepth(root *TreeNode) int {
	ret := 0
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 1
	for len(queue) != 0 {
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count = len(queue)
		ret++
	}
	return ret
}