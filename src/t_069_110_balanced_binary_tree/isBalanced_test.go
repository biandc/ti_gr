package t_069_110_balanced_binary_tree

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var (
		root     = []int{3, 9, 20, 0, 0, 15, 7}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = true
		actual   bool
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &TreeNode{Val: val}
		}
	}
	for index, val := range queue {
		if val == nil {
			continue
		}
		i := index * 2
		if i+1 < l && queue[i+1] != nil {
			val.Left = queue[i+1]
		}
		if i+2 < l && queue[i+2] != nil {
			val.Right = queue[i+2]
		}
	}
	actual = isBalanced(queue[0])
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}

}
