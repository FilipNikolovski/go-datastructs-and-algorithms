package trie_test

import (
	"errors"
	"testing"

	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/trees/trie"
)

func TestTrie(t *testing.T) {
	tr := trie.NewTrie()

	words := []string{"doggy", "dog", "dot", "hug", "hugo", "kate"}

	for _, w := range words {
		tr.Put(w)
		if !tr.Find(w) {
			t.Errorf("unable to find %s", w)
		}
	}

	if tr.Find("dogg") {
		t.Error("was able to find dog even though it's not one of the words in the trie")
	}

	for _, w := range words {
		err := tr.Delete(w)
		if err != nil {
			t.Errorf("unable to delete %s: %s", w, err)
			continue
		}
		if tr.Find(w) {
			t.Errorf("was able to find '%s' even though it's deleted", w)
		}
	}

	err := tr.Delete("unknown")
	if !errors.Is(err, trie.ErrWordNotFound) {
		t.Errorf("expected error to be %s but error is %s", trie.ErrWordNotFound, err)
	}
}
