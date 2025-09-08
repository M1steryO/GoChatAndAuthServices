package main

import "fmt"

var res int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	res = max(res, left+right)
	fmt.Println(res)
	return max(left, right) + 1

}
func diameterOfBinaryTree(root *TreeNode) int {
	helper(root)
	return res
}
func main() {
	diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	})
}
