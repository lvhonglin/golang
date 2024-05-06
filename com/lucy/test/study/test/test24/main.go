package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 1)
	queue[0] = root
	for {
		if len(queue) == 0 {
			break
		}
		oldLen := len(queue)
		var i int
		var pre *Node
		for ; i < oldLen; i++ {
			tmp := queue[i]
			if pre == nil {
				pre = tmp
			} else {
				pre.Next = tmp
				pre = tmp
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		queue = queue[oldLen:]
	}
	return root

}

func main() {
	node1 := &Node{
		Val: 1,
	}
	node2 := &Node{
		Val: 2,
	}
	node3 := &Node{
		Val: 3,
	}
	node4 := &Node{
		Val: 4,
	}
	node5 := &Node{
		Val: 5,
	}
	node6 := &Node{
		Val: 6,
	}
	node7 := &Node{
		Val: 7,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node := connect(node1)
	println(node)
}
