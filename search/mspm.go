// Package mspm provides model that will have collection of trieNodes that represent the patterns to be searched in a document.
package mspm

import (
	"bufio"
	"io"
	"io/ioutil"

	"github.com/BlackRabbitt/mspm/ds/trie"
)

// M represents mspm-model.
type M struct {
	Name     string // Name representing mspm model
	trieNode *trie.TrieHashNode
}

// Output defines the output of mspm.
// string - Term found in document
// int32 - Count of term in that document
type Output map[string]int32

// NewModel will return a fresh new model
func NewModel(name string) *M {
	return &M{Name: name, trieNode: trie.NewTrieHashNode()}
}

// Build trie datastructure that accepts multiline list of words.
func (model *M) Build(words io.Reader) {
	scanner := bufio.NewScanner(words)
	for scanner.Scan() {
		model.trieNode.Insert(scanner.Bytes())
	}
}

// MultiTermMatch returns all the trie-terms found in document.
func (model *M) MultiTermMatch(document io.Reader) (output Output, err error) {
	output = make(map[string]int32)
	content, err := ioutil.ReadAll(document)

	if err != nil {
		return
	}

	var index byte
	tNode := model.trieNode

	// start and end pointer select the current valid term. It is adjusted itself over time.
	startPointer := 0
	endPointer := startPointer

	for level := startPointer; level < len(content); level++ {
		index = content[level]

		if tNode.Children[index] == nil {
			if endPointer > startPointer {
				term := string(content[startPointer : endPointer+1])
				output[term]++
			}
			startPointer = level + 1
			endPointer = startPointer

			tNode = model.trieNode
			continue
		}

		tNode = tNode.Children[index]

		if tNode.FinalState {
			endPointer = level
			if len(tNode.Children) > 0 {
				continue
			}
			term := string(content[startPointer : endPointer+1]) // exclusive
			output[term]++
			startPointer = level + 1
			tNode = model.trieNode
			continue
		}
	}
	return
}
