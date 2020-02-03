package app

import (
	"basketball/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-collections/collections/trie"
	"github.com/meirf/gopart"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func fetchData(t *trie.Trie)error{
	if err := fetchFromSource(t); err != nil {
		msg := "unable to pull data from source api. %v"
		log.Println(fmt.Sprintf(msg, err))
		return fetchFromFile(t)
	}
	return nil
}

func fetchFromSource(t *trie.Trie)error{
	return errors.New("dummy error")
}

func fetchFromS3(){
	//TODO setup S3 integration for storing daily player data as fallback
}

func fetchFromFile(t *trie.Trie)error{
	//TODO remove once S3 integration complete
	var players = make([]model.Player, 0)

	//open json file of all players
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	//slurp bytes
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	//unmarshall into array
	err = json.Unmarshal(byteValue, &players)
	if err != nil {
		return err
	}

	//hash and store in trie
	for indexRange := range gopart.Partition(len(players), 10){
		go save(players[indexRange.Low:indexRange.High], t)
	}

	return nil
}

func ratePlayers(config model.ScoreConfig, t *trie.Trie)[]model.Player{
	return scorePlayers(config, t)
}

func save(players []model.Player, t *trie.Trie){
	for _, player := range players {
		now := time.Now()
		player.CreatedDateTime = now
		player.UpdatedDateTime = now
		key, err := hash(player)
		if err != nil{
			log.Println("could not hash player")
			log.Println(err)
		}
		storePlayer(player)
		t.Insert(key, player)
	}
}
