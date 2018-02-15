package sei

type Node struct {
	value    rune
	data     interface{}
	children map[rune]*Node
	isLeaf   bool
}

func (n *Node) Data() interface{} {
	return n.data
}

func (n *Node) IsLeaf() bool {
	return n.isLeaf
}

func NewNode(val rune, isLeaf bool) *Node {
	return &Node{
		value:    val,
		children: make(map[rune]*Node),
		isLeaf:   isLeaf,
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(0, false),
	}
}

func (t *Trie) Add(key string, data interface{}) *Node {
	runes := []rune(key)
	node := t.root

	for _, r := range runes {
		if _, ok := node.children[r]; !ok {
			node.children[r] = NewNode(r, false)
		}
		node = node.children[r]
	}

	node.isLeaf = true
	node.data = data
	return node
}

func (t *Trie) Find(key string) (*Node, bool) {
	return t.find(t.root, key)
}

func (t *Trie) find(node *Node, key string) (*Node, bool) {
	runes := []rune(key)

	if node == nil || len(runes) == 0 {
		return nil, false
	}

	for _, r := range runes {
		n, ok := node.children[r]
		if ok {
			node = n
		} else {
			return nil, false
		}
	}

	return node, true
}
