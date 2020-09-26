package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	left.Right = &right
	res := isBalanced(&root)
	fmt.Println(res)
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := dfs(root)
	return isBalanced
}

func dfs(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := dfs(root.Left)
	if leftIsBalanced == false {
		return 0, false
	}
	rightDepth, rightIsBalanced := dfs(root.Right)
	if rightIsBalanced == false {
		return 0, false
	}

	if -1 <= leftDepth-rightDepth &&
		leftDepth-rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
