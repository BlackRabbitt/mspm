// mspm - Multi String Pattern-Matching
// Information Retrival System
// mspm implements trieHashNode data structure for building trie for list of words.
// specific/multiple keyword can then be searched in the tree in linear time.
// refer to example for more detail on implementation
package mspm

import (
	"bufio"
	"io"
	"io/ioutil"

	"github.com/BlackRabbitt/mspm/ds/trie"
)

// mspm Model.
type M struct {
	Name     string // Name representing mspm model
	trieNode *trie.TrieHashNode
}

// Output defines the output of mspm.
// string - Term found in document
// int32 - Count of term in that document
type Output map[string]int32

func NewModel(name string) *M {
	return &M{Name: name, trieNode: trie.NewTrieHashNode()}
}

// Build builds trie datastructure that accepts multiline list of words.
func (self *M) Build(words io.Reader) {
	scanner := bufio.NewScanner(words)
	for scanner.Scan() {
		self.trieNode.Insert(scanner.Bytes())
	}
}

// Returns all the trie-terms found in document.
func (self *M) MultiTermMatch(document io.Reader) (output Output, err error) {
	output = make(map[string]int32)
	content, err := ioutil.ReadAll(document)

	if err != nil {
		return
	}

	var index byte
	tNode := self.trieNode

	// start and end pointer select the current valid term. It is adjusted itself over time.
	startPointer := 0
	endPointer := startPointer

	for level := startPointer; level < len(content); level++ {
		index = content[level]

		if tNode.Children[index] == nil {
			if endPointer > startPointer {
				term := string(content[startPointer : endPointer+1])
				output[term] += 1
			}
			startPointer = level + 1
			endPointer = startPointer

			tNode = self.trieNode
			continue
		}

		tNode = tNode.Children[index]

		if tNode.FinalState {
			endPointer = level
			if len(tNode.Children) > 0 {
				continue
			}
			term := string(content[startPointer : endPointer+1]) // exclusive
			output[term] += 1
			startPointer = level + 1
			tNode = self.trieNode
			continue
		}
	}
	return
}
