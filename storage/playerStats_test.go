package storage

import (
	"basketball/model"
	"github.com/golang-collections/collections/trie"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestSave(t *testing.T) {
	players := []model.Player{
		model.Player{Id: 0, Name: "A"},
		model.Player{Id: 1, Name: "B"},
	}

	testTrie := trie.New()
	testing.Init()

	save(players, testTrie)

	assert.Equal(t, len(players), testTrie.Len())
	testTrie.Do(func(k, v interface{}) bool {
		p := v.(model.Player)
		assert.Equal(t, false, p.CreatedDateTime.IsZero())
		assert.Equal(t, false, p.UpdatedDateTime.IsZero())
		return true
	})
}
