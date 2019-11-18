package main

import "fmt"

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	l := len(preorder)
	if l == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if l == 1 {
		return root
	}
	leftLen := 0
	rightLen := 0
	for _, v := range inorder {
		if v == root.Val {
			break
		}
		leftLen++
	}
	rightLen = l - leftLen - 1
	if leftLen > 0 {
		root.Left = buildTree(preorder[1:leftLen+1], inorder[0:leftLen])
	}
	if rightLen > 0 {
		root.Right = buildTree(preorder[leftLen+1:], inorder[leftLen+1:])
	}
	return root
}

func main() {
	pre:=[]int{3,9,20,15,7}
	in:=[]int{9,3,15,20,7}
	tree := buildTree(pre, in)
	fmt.Print(tree)
}
