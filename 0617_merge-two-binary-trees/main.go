package main

import "fmt"

func main() {
	first := TreeNode{Val:1}
	second := TreeNode{Val:2}
	third := TreeNode{Val:4}
	secondRight := TreeNode{Val:3}

	first.Left = &second
	first.Right = &secondRight
	second.Left = &third


	first1 := TreeNode{Val:10}
	second1 := TreeNode{Val:20}
	third1 := TreeNode{Val:40}
	secondRight1 := TreeNode{Val:30}

	first1.Left = &second1
	first1.Right = &secondRight1
	second1.Left = &third1

	fmt.Println(mergeTrees(&first,&first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil{
		return nil
	}

	if t1 == nil{
		return t2
	}

	if t2 == nil{
		return t1
	}

	root := &TreeNode{Val:t1.Val+t2.Val}

	root.Left = mergeTrees(t1.Left,t2.Left)
	root.Right = mergeTrees(t1.Right,t2.Right)

	return root
}
