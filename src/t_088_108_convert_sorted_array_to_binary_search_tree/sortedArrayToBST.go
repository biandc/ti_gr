package t_088_108_convert_sorted_array_to_binary_search_tree

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

func sortedArrayToBST(nums []int) *TreeNode {
	var createNode func(nums []int) *TreeNode
	createNode = func(nums []int) *TreeNode {
		l := len(nums)
		if l == 0 {
			return nil
		}
		mid := l / 2
		node := &TreeNode{
			Val: nums[mid],
		}
		node.Left = createNode(nums[:mid])
		node.Right = createNode(nums[mid+1:])
		return node
	}
	return createNode(nums)
}
