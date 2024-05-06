package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var tmp *TreeNode
	prep := &tmp
	dfs(root, prep)
}

func dfs(root *TreeNode, prep **TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right, prep)
	dfs(root.Left, prep)
	if *prep == nil {
		*prep = root
	} else {
		root.Right = *prep
		root.Left = nil
		*prep = root
	}
}
func main() {
	var tmp *TreeNode
	prep := &tmp
	println(*prep)
}
