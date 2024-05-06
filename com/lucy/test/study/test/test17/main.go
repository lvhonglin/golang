package main

func main() {
	queue := NewMyQueue()
	queue.offer(&TreeNode{Val: 1})
	queue.offer(&TreeNode{Val: 2})
	queue.offer(&TreeNode{Val: 3})
	println(queue.take().Val)
	println(queue.take().Val)
	println(queue.take().Val)
}
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	mQueue := NewMyQueue()
	mQueue.offer(root)
	for {
		if mQueue.size() == 0 {
			break
		}
		tmp := make([]int, mQueue.size())
		tmpSize := mQueue.size()
		for {
			if tmpSize == 0 {
				break
			}
			tNode := mQueue.take()
			tmp = append(tmp, tNode.Val)
			if tNode.Left != nil {
				mQueue.offer(tNode.Left)
			}
			if tNode.Right != nil {
				mQueue.offer(tNode.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MQueue struct {
	last  *Node
	first *Node
	sizeQ int
}
type Node struct {
	treeNode *TreeNode
	next     *Node
}

func NewMyQueue() *MQueue {
	return &MQueue{}
}
func (m *MQueue) size() int {
	return m.sizeQ
}
func (m *MQueue) offer(node *TreeNode) {
	tmp := &Node{
		treeNode: node,
	}
	if m.first == nil {
		m.first = tmp
		m.last = tmp
	} else {
		m.last.next = tmp
		m.last = tmp
	}
	m.sizeQ++
}
func (m *MQueue) take() *TreeNode {
	tmp := m.first.next
	res := m.first.treeNode
	if m.first == m.last {
		m.last = tmp
	}
	m.first = tmp
	return res
}
