package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)

	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
