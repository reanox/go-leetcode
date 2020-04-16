package binary_tree_preorder_traversal

import . "github.com/reanox/go-leetcode/structures"

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	helper(root.Left, res)
	helper(root.Right, res)
}
