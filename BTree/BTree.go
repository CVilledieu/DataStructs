package main

/*
Rules of BTree
all leaf nodes have the same depth
each node has n elements where min > 0 and min <= n <= 2 * min
each non root node has n elements and n + 1 sub trees
root may have fewer than min elements
*/

type Node struct {
	Leaf     bool
	Key      []int
	Children []*Node
	Data     string
}

type BTree struct {
	Root     *Node
	MaxNodes int
	MinNodes int
}

func newNode(key int) *Node {
	return &Node{Key: []int{key}}
}

func (n *Node) addKeyToNode(newKey int) {
	n.Key = append(n.Key, newKey)
}

func (n *Node) addChild(c *Node) {
	n.Children = append(n.Children, c)
}

func createTree(size, key int) *BTree {
	root := newNode(key)
	tree := BTree{Root: root}
	tree.MinNodes = size
	tree.MaxNodes = size * 2
	return &tree
}

func (n *Node) search(key int) *Node {
	for _, node := range n.Key {
		if node == key {
			return n
		}
	}

	return
}

func main() {

}
