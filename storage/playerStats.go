package storage

import (
	"basketball/client"
	"basketball/env"
	"basketball/model"
	"context"
	"fmt"
	"github.com/golang-collections/collections/trie"
	"github.com/meirf/gopart"
	"github.com/slham/toolbelt/l"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var store *trie.Trie

func Store() *trie.Trie {
	return store
}

func Initialize(config env.Config) bool {
	store = trie.New()
	store.Init()

	switch config.Env {
	case "dev":
		return fetchFromLocal(store, config.Storage.FileName)
	case "prod":
		return fetchFromS3(store, config.Storage.Bucket, config.Storage.Prefix)
	default:
		l.Error(nil, "invalid environment configuration")
		return false
	}
}

func fetchFromLocal(t *trie.Trie, fileName string) bool {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "basketball") {
		wd = filepath.Dir(wd)
	}
	path := fmt.Sprintf("%s/%s", wd, fileName)
	//envPath, _ := filepath.Abs(path)
	l.Debug(nil, "path:%s", path)
	playersBytes, err := ioutil.ReadFile(path)
	if err != nil {
		l.Error(nil, "unable to read player stats from local file: %s", path)
		return false
	}

	err = UnmarshalAndSavePlayers(nil, playersBytes, t)
	if err != nil {
		return false
	}

	return true
}

func fetchFromS3(t *trie.Trie, bucket, prefix string) bool {
	client.InitializeSession()

	key, err := client.GetLatestS3Key(bucket, prefix)
	if err != nil {
		return false
	}

	playersBytes, err := client.GetS3Object(bucket, key)
	if err != nil {
		return false
	}

	err = UnmarshalAndSavePlayers(nil, playersBytes, t)
	if err != nil {
		return false
	}

	return true
}

func UnmarshalAndSavePlayers(ctx context.Context, playersBytes []byte, t *trie.Trie) error {
	players := make([]model.Player, 0)
	err := yaml.Unmarshal(playersBytes, &players)
	if err != nil {
		l.Error(ctx, "unable to convert players: %v", err)
		return err
	}

	partitionSave(ctx, players, t)
	return nil
}

func partitionSave(ctx context.Context, players []model.Player, t *trie.Trie) {
	var wg sync.WaitGroup
	for indexRange := range gopart.Partition(len(players), 10) {
		wg.Add(1)
		go save(ctx, players[indexRange.Low:indexRange.High], t, &wg)
	}
	wg.Wait()
}

func save(ctx context.Context, players []model.Player, t *trie.Trie, wg *sync.WaitGroup) {
	defer wg.Done()
	l.Debug(ctx, "storing players")
	for _, player := range players {
		now := time.Now()
		player.CreatedDateTime = now
		player.UpdatedDateTime = now
		key, err := hash(player)
		if err != nil {
			l.Error(ctx, "could not hash player: %v", err)
			continue
		}
		t.Insert(key, player)
	}
}
