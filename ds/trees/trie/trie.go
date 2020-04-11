package trie

import (
	"errors"

	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/stacks/arraystack"
)

type Trie struct {
	children map[rune]*Trie
	end      bool
}

func NewTrie() *Trie {
	return &Trie{}
}

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
	ErrWordNotFound = errors.New("trie: word not found")
	ErrAssertType   = errors.New("trie: type assertion: not a nodeRune")
)

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
