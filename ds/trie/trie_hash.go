package trie

// HashNode represents TrieNode using hash-table
type HashNode struct {
	// Child nodes
	Children map[byte]*HashNode

	// True if current node reached final state
	FinalState bool
}

// NewHashNode returns new trie hashNode
func NewHashNode() (tNode *HashNode) {
	tNode = &HashNode{Children: make(map[byte]*HashNode), FinalState: false}
	return
}

// Insert key in trie datastructure
func (root *HashNode) Insert(key []byte) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			tNode.Children[index] = NewHashNode()
		}
		tNode = tNode.Children[index]
	}

	// mark last node as final node.
	tNode.FinalState = true
}

// Search for key in trie datastructure
// returns found=true if the given string is found and returns final=true if the string reached final state in trie state machine
func (root *HashNode) Search(key []byte) (found bool, final bool) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			return false, false
		}
		tNode = tNode.Children[index]
	}
	return tNode != nil, tNode.FinalState
}
