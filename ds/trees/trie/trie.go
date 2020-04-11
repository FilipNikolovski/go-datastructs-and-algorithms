package trie

import (
	"errors"

	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/stacks/arraystack"
)

// Trie is a trie of runes. Each trie node has an 'end' bool flag, which
// indicates whether the node represents the end character of a word.
type Trie struct {
	children map[rune]*Trie
	end      bool
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{}
}

// Find walks the trie by the given string and checks if the last node represents
// an ending of a word or not.
func (t *Trie) Find(word string) bool {
	node := t
	for _, r := range word {
		node = node.children[r]
		if node == nil {
			return false
		}
	}
	return node.end
}

// Put adds a new word to the trie.
func (t *Trie) Put(word string) {
	node := t
	for _, r := range word {
		if node.children[r] == nil {
			if node.children == nil {
				node.children = make(map[rune]*Trie)
			}
			node.children[r] = &Trie{}
		}
		node = node.children[r]
	}
	node.end = true
}

type nodeRune struct {
	r    rune
	node *Trie
}

var (
	// ErrWordNotFound is returned when a word is not found.
	ErrWordNotFound = errors.New("trie: word not found")
	// ErrAssertType is returned if the type is not a nodeRune.
	ErrAssertType = errors.New("trie: type assertion: not a nodeRune")
)

// Delete deletes a word from the trie.
func (t *Trie) Delete(word string) error {
	s := arraystack.New()
	node := t
	for _, r := range word {
		s.Push(&nodeRune{r: r, node: node})
		node = node.children[r]
		if node == nil {
			return ErrWordNotFound
		}
	}

	// no longer indicating an end of a word
	node.end = false

	if !node.isLeafNode() { // parent has other nodes, we stop here.
		return nil
	}

	for !s.Empty() {
		el, err := s.Pop()
		if err != nil {
			return err
		}
		nr, ok := el.(*nodeRune)
		if !ok {
			return ErrAssertType
		}

		delete(nr.node.children, nr.r)

		if !nr.node.isLeafNode() {
			break
		}

		nr.node.children = nil

		if nr.node.end {
			break
		}
	}

	return nil
}

func (t *Trie) isLeafNode() bool {
	return len(t.children) == 0
}
