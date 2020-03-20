package storage

import (
	"basketball/client"
	"basketball/env"
	"basketball/model"
	"context"
	"github.com/golang-collections/collections/trie"
	"github.com/meirf/gopart"
	"github.com/slham/toolbelt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
		toolbelt.Error(nil, "invalid environment configuration")
		return false
	}
}

func fetchFromLocal(t *trie.Trie, fileName string) bool {
	playersBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		toolbelt.Error(nil, "unable to read player stats from local file: %s", fileName)
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
		toolbelt.Error(ctx, "unable to convert players: %v", err)
		return err
	}

	c := make(chan bool)
	go partitionSave(ctx, c, players, t)
	<-c
	return nil
}

func partitionSave(ctx context.Context, c chan bool, players []model.Player, t *trie.Trie) {
	for indexRange := range gopart.Partition(len(players), 10) {
		go save(ctx, players[indexRange.Low:indexRange.High], t)
	}
	c <- true
}

func save(ctx context.Context, players []model.Player, t *trie.Trie) {
	toolbelt.Debug(ctx, "storing players")
	for _, player := range players {
		now := time.Now()
		player.CreatedDateTime = now
		player.UpdatedDateTime = now
		key, err := hash(player)
		if err != nil {
			toolbelt.Error(ctx, "could not hash player: %v", err)
		}
		t.Insert(key, player)
	}
}
