package storage

import (
	"basketball/model"
	"github.com/golang-collections/collections/trie"
	"gopkg.in/go-playground/assert.v1"
	"sync"
	"testing"
)

func TestFetchFromLocal(t *testing.T) {
	ids := []int{
		20000439,
		20000441,
		20000442,
		20000443,
		20000452,
		20000453,
		20000455,
		20000456,
		20000457,
		20000459,
		20000460,
		20000464,
		20000466,
		20000468,
		20000471,
		20000474,
		20000475,
		20000482,
		20000483,
	}

	testTrie := trie.New()
	testTrie.Init()
	ok := fetchFromLocal(testTrie, "1583510437-test.yaml")
	assert.Equal(t, ok, true)
	assert.Equal(t, len(ids), testTrie.Len())
	testTrie.Do(func(k,v interface{}) bool{
		p := v.(model.Player)
		found := false
		for _, id := range ids {
			if p.Id == id{
				found = true
			}
		}
		return found
	})
}

func TestFetchFromLocalNoFile(t *testing.T){
	testTrie := trie.New()
	testTrie.Init()
	ok := fetchFromLocal(testTrie, "blah")
	assert.Equal(t, ok, false)
}

func TestSave(t *testing.T) {
	players := []model.Player{
		model.Player{Id: 0, Name: "A"},
		model.Player{Id: 1, Name: "B"},
	}

	testTrie := trie.New()
	testing.Init()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t.Log("test routine")
	}()

	save(nil, players, testTrie, &wg)

	assert.Equal(t, len(players), testTrie.Len())
	testTrie.Do(func(k, v interface{}) bool {
		p := v.(model.Player)
		assert.Equal(t, false, p.CreatedDateTime.IsZero())
		assert.Equal(t, false, p.UpdatedDateTime.IsZero())
		return true
	})
}
