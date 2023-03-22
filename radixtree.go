package httprouter

// RadixTreeNode represents a node in the radix tree
type RadixTreeNode struct {
	Key      string
	Value    interface{}
	Children map[string]*RadixTreeNode
}

// NewRadixTreeNode creates a new radix tree node
func NewRadixTreeNode(key string, value interface{}) *RadixTreeNode {
	return &RadixTreeNode{
		Key:      key,
		Value:    value,
		Children: make(map[string]*RadixTreeNode),
	}
}

// RadixTree represents the radix tree
type RadixTree struct {
	Root *RadixTreeNode
}

// NewRadixTree creates a new radix tree
func NewRadixTree() *RadixTree {
	return &RadixTree{
		Root: NewRadixTreeNode("", nil),
	}
}

// Insert inserts a key-value pair into the radix tree
func (tree *RadixTree) Insert(key string, value interface{}) {
	currentNode := tree.Root
	for {
		matchLength := longestCommonPrefixLength(currentNode.Key, key)
		if matchLength == len(currentNode.Key) {
			key = key[matchLength:]
			if len(key) == 0 {
				currentNode.Value = value
				return
			}
			if child, ok := currentNode.Children[key[0:1]]; ok {
				currentNode = child
			} else {
				currentNode.Children[key[0:1]] = NewRadixTreeNode(key, value)
				return
			}
		} else {
			newNode := NewRadixTreeNode(currentNode.Key[matchLength:], currentNode.Value)
			newNode.Children = currentNode.Children
			currentNode.Key = currentNode.Key[:matchLength]
			currentNode.Value = nil
			currentNode.Children = make(map[string]*RadixTreeNode)
			currentNode.Children[newNode.Key[0:1]] = newNode
			if len(key) > matchLength {
				key = key[matchLength:]
				currentNode.Children[key[0:1]] = NewRadixTreeNode(key, value)
			} else {
				currentNode.Value = value
			}
			return
		}
	}
}

// Search finds a key in the radix tree and returns its value
func (tree *RadixTree) Search(key string) (interface{}, bool) {
	currentNode := tree.Root
	for {
		matchLength := longestCommonPrefixLength(currentNode.Key, key)
		if matchLength == len(currentNode.Key) {
			key = key[matchLength:]
			if len(key) == 0 {
				return currentNode.Value, true
			}
			if child, ok := currentNode.Children[key[0:1]]; ok {
				currentNode = child
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	}
}

// longestCommonPrefixLength returns the length of the longest common prefix between two strings
func longestCommonPrefixLength(s1, s2 string) int {
	minLength := min(len(s1), len(s2))
	for i := 0; i < minLength; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return minLength
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
