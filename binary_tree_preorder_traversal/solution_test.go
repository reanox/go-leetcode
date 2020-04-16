package binary_tree_preorder_traversal

import (
	"testing"

	. "github.com/reanox/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
)

// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// Output: [1,2,3]
func Test_preorder_traversal(t *testing.T) {

	testName := "Preorder traversal"
	input := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Right: nil,
				Left:  nil,
			},
			Right: nil,
		},
	}
	output := []int{1, 2, 3}

	t.Run(testName, func(t *testing.T) {
		assert.Equal(t, preorderTraversal(&input), output)
	})
}
