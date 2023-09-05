package algo

type Node struct {
	Val          int
	Next, Random *Node
}

// https://leetcode.com/problems/copy-list-with-random-pointer
func copyRandomList(head *Node) *Node {
	mp, node := make(map[*Node]*Node), head
	for node != nil {
		mp[node] = &Node{Val: node.Val}
		node = node.Next
	}
	node = head
	for node != nil {
		mp[node].Next = mp[node.Next]
		mp[node].Random = mp[node.Random]
		node = node.Next
	}
	return mp[head]
}
