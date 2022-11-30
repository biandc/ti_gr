package t_063_590_n_ary_tree_postorder_traversal

import (
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	var (
		root      = []int{1, 0, 3, 2, 4, 0, 5, 6, 0}
		queue     = make([]*Node, len(root))
		nodesList = [][]*Node{}
		zeroList  = []int{}
		expected  = []int{5, 6, 3, 2, 4, 1}
		actual    []int
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &Node{Val: val}
		} else {
			zeroList = append(zeroList, index)
		}
	}
	start := 0
	for _, val := range zeroList {
		nodesList = append(nodesList, queue[start:val])
		start = val + 1
	}
	i := 1
	for _, val := range queue {
		if val != nil {
			val.Children = nodesList[i]
			i++
			if i >= len(nodesList) {
				break
			}
		}
	}
	actual = postorder(queue[0])
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}

}
