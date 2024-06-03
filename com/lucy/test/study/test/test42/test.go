package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	m := make(map[*Node]*Node)
	return dfs(node, m)
}

func dfs(node *Node, mm map[*Node]*Node) *Node {

	n, ok := mm[node]
	if ok {
		return n
	}
	tmp := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	if node.Neighbors != nil && len(node.Neighbors) != 0 {
		for _, v := range node.Neighbors {
			tmp.Neighbors = append(tmp.Neighbors, dfs(v, mm))
		}
	}
	mm[node] = tmp
	return tmp

}
