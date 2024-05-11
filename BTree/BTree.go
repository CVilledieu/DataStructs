package main

import "os"

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
}

type BTree struct {
	Root     *Node
	MaxNodes int
	MinNodes int
}

func newNode(key int) *Node {
	return &Node{Key: []int{key}, Leaf: true}
}

func (n *Node) addKeyToNode(newKey int) {
	n.Key = append(n.Key, newKey)
}

func (n *Node) addChild(c *Node) {
	n.Children = append(n.Children, c)
}

func createTree(size, keys int) *BTree {
	root := newNode(key)
	tree := BTree{Root: root}
	tree.MinNodes = size
	tree.MaxNodes = size * 2
	return &tree
}

func (n *Node) addToTree(key int) {
	if n.Key[0] > key {
		if n.Children[0].Leaf == true{
			newNode := newNode(key)
			newNode.Leaf = true
			n.addChild(newNode)
		}
		n.Children[0].addToTree(key)
	}
	for i, ck := range n.Key {
		if ck <= key {
			n.Children[i].addToTree(key)
		}
	}
}


func (t *BTree) insert(n *Node) {

}

//will return closest none leaf node
func (n *Node) search(key int) (*Node, bool) {
	//cnk is current key
	for i, ck := n.Key {

	}
}

func main() {

}




func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_,err = fp.Write(data)
	if err != nil{
		return err
	}
	return fp.Sync()
}