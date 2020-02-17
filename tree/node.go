package tree

type Address string

type Tree struct {
	root *Node
	pair map[*Node]*Node
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
		pair: map[*Node]*Node{},
	}
}

func (tree *Tree) AddNode(child *Node, parent *Node) {
	if parent == nil {
		tree.root = child
	} else {
		tree.pair[child] = parent
	}
}

func (tree *Tree) DeleteNode(node *Node) []*Node {
	var list []*Node

	delete(tree.pair, node)
	list = append(list, node)

	for c, p :=range tree.pair {
		if p == node {
			delete(tree.pair, c)
			list = append(list, c)

			ls := tree.DeleteNode(c)
			list = append(list, ls...)
		}
	}
	return list
}

func (tree *Tree) ChildNodes(node *Node) []*Node {
	var list []*Node
	for c, p := range tree.pair {
		if p == node {
			list = append(list, c)
			l := tree.ChildNodes(c)
			list = append(list, l...)
		}
	}

	return list
}

type Node struct {
	GroupID string  `json:"group_id"`
	Address Address `json:"address"`
	Parent  *Node   `json:"parent"`
	Content string  `json:"content"`
	Level   int     `json:"level"`
}

func NewNode(groupId string, addr Address, Content string, parent *Node) *Node {
	var node Node
	node.GroupID = groupId
	node.Address = addr
	node.Content = Content
	if parent == nil {
		node.Parent = nil
		node.Level = 0
	} else {
		node.Parent = parent
		node.Level = parent.Level + 1
	}
	return &node
}

