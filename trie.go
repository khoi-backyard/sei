package sei

import (
	"net/http"
)

type (
	node struct {
		value         rune
		children      map[rune]*node
		isLeaf        bool
		methodHandler *methodHandler
	}

	methodHandler struct {
		get     HandlerFunc
		head    HandlerFunc
		post    HandlerFunc
		put     HandlerFunc
		patch   HandlerFunc
		delete  HandlerFunc
		connect HandlerFunc
		options HandlerFunc
		trace   HandlerFunc
	}
)

func (n *node) setHandler(handler HandlerFunc, methods ...string) {
	for _, m := range methods {
		switch m {
		case http.MethodGet:
			n.methodHandler.get = handler
		case http.MethodHead:
			n.methodHandler.head = handler
		case http.MethodPost:
			n.methodHandler.post = handler
		case http.MethodPut:
			n.methodHandler.put = handler
		case http.MethodPatch:
			n.methodHandler.patch = handler
		case http.MethodDelete:
			n.methodHandler.delete = handler
		case http.MethodConnect:
			n.methodHandler.connect = handler
		case http.MethodOptions:
			n.methodHandler.options = handler
		case http.MethodTrace:
			n.methodHandler.trace = handler
		}
	}
}

func (n *node) getHandler(method string) HandlerFunc {
	if !n.isLeaf {
		return nil
	}
	switch method {
	case http.MethodGet:
		return n.methodHandler.get
	case http.MethodHead:
		return n.methodHandler.head
	case http.MethodPost:
		return n.methodHandler.post
	case http.MethodPut:
		return n.methodHandler.put
	case http.MethodPatch:
		return n.methodHandler.patch
	case http.MethodDelete:
		return n.methodHandler.delete
	case http.MethodConnect:
		return n.methodHandler.connect
	case http.MethodOptions:
		return n.methodHandler.options
	case http.MethodTrace:
		return n.methodHandler.trace
	default:
		return nil
	}
}

func NewNode(val rune, isLeaf bool) *node {
	return &node{
		value:         val,
		children:      make(map[rune]*node),
		isLeaf:        isLeaf,
		methodHandler: new(methodHandler),
	}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(0, false),
	}
}

func (t *Trie) Add(handler HandlerFunc, path string, methods ...string) {
	runes := []rune(path)
	node := t.root

	for _, r := range runes {
		if _, ok := node.children[r]; !ok {
			node.children[r] = NewNode(r, false)
		}
		node = node.children[r]
	}

	node.isLeaf = true
	node.setHandler(handler, methods...)
}

func (t *Trie) Find(path, method string) HandlerFunc {
	return t.find(t.root, path, method)
}

func (t *Trie) find(node *node, path, method string) HandlerFunc {
	runes := []rune(path)

	if node == nil || len(runes) == 0 {
		return nil
	}

	for _, r := range runes {
		n, ok := node.children[r]
		if ok {
			node = n
		} else {
			return nil
		}
	}

	return node.getHandler(method)
}
