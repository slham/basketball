package storage

import (
	"basketball/client"
	"basketball/env"
	"basketball/model"
	"github.com/golang-collections/collections/trie"
	"github.com/meirf/gopart"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
		log.Println("error loading ENVIRONMENT env variable")
		return false
	}
}

func fetchFromLocal(t *trie.Trie, fileName string) bool {
	playersBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("unable to read player stats from local file")
		log.Println(fileName)
		log.Println(err)
		return false
	}

	err = UnmarshalAndSavePlayers(playersBytes, t)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func fetchFromS3(t *trie.Trie, bucket, prefix string) bool {
	client.InitializeSession()

	key, err := client.GetLatestS3Key(bucket, prefix)
	if err != nil {
		log.Println("unable to find latest S3 Object Key")
		log.Println(err)
		return false
	}

	playersBytes, err := client.GetS3Object(bucket, key)
	if err != nil {
		log.Println("unable to read player stats from S3")
		log.Println(err)
		return false
	}

	err = UnmarshalAndSavePlayers(playersBytes, t)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func UnmarshalAndSavePlayers(playersBytes []byte, t *trie.Trie) error {
	players := make([]model.Player, 0)
	err := yaml.Unmarshal(playersBytes, &players)
	if err != nil {
		log.Println("unable to convert players")
		return err
	}

	c := make(chan bool)
	go partitionSave(c, players, t)
	<-c
	return nil
}

func partitionSave(c chan bool, players []model.Player, t *trie.Trie) {
	for indexRange := range gopart.Partition(len(players), 10) {
		go save(players[indexRange.Low:indexRange.High], t)
	}
	c <- true
}

func save(players []model.Player, t *trie.Trie) {
	log.Println("storing players")
	for _, player := range players {
		now := time.Now()
		player.CreatedDateTime = now
		player.UpdatedDateTime = now
		key, err := hash(player)
		if err != nil {
			log.Println("could not hash player")
			log.Println(err)
		}
		t.Insert(key, player)
	}
}
