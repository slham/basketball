package app

import (
	"basketball/client"
	"basketball/model"
	"errors"
	"github.com/golang-collections/collections/trie"
	"github.com/meirf/gopart"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func (a *App) ratePlayers(config model.ScoreConfig) []model.Player {
	return scorePlayers(config, a.store)
}

func fetchData(t *trie.Trie) error {
	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "DEV":
		return fetchFromLocal(t)
	case "PROD":
		return fetchFromS3(t)
	default:
		return errors.New("error loading ENVIRONMENT env variable")
	}
}

func fetchFromLocal(t *trie.Trie) error {
	playersBytes, err := ioutil.ReadFile("1583510437.yaml")
	if err != nil {
		log.Println("unable to read player stats from local file")
		return err
	}

	err = unmarshalAndSavePlayers(playersBytes, t)
	if err != nil {
		return err
	}

	return nil
}

func fetchFromS3(t *trie.Trie) error {
	client.InitializeSession()

	key, err := client.GetLatestS3Key("sheldonsandbox-basketball", "player-stats/2020")
	if err != nil {
		log.Println("unable to find latest S3 Object Key")
		return err
	}

	playersBytes, err := client.GetS3Object("sheldonsandbox-basketball", key)
	if err != nil {
		log.Println("unable to read player stats from S3")
		return err
	}

	err = unmarshalAndSavePlayers(playersBytes, t)
	if err != nil {
		return err
	}

	return nil
}

func unmarshalAndSavePlayers(playersBytes []byte, t *trie.Trie) error {
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
