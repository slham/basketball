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
	"net/http"
	"os"
	"time"
)

func fetchData(t *trie.Trie) error {
	return fetchFromSource(t)
}

//TODO: implement retry logic
//#downstream_pull
func fetchFromSource(t *trie.Trie) error {
	key := os.Getenv("NBA_API_KEY")
	if key == "" {
		return errors.New("unable to load environment variables")
	}

	url := fmt.Sprintf("https://api.sportsdata.io/v3/nba/stats/json/PlayerSeasonStats/2020?key=%s", key)
	res, err := http.Get(url)
	if err != nil {
		log.Println(fmt.Sprintf("error sending to %v", url))
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("unable to read player stats response")
		return err
	}

	players := make([]model.Player, 0)

	err = json.Unmarshal(body, &players)
	if err != nil {
		log.Println(string(body))
		log.Println("unable to convert players")
		return err
	}

	//#async
	c := make(chan bool)
	go func() {
		//hash and store in trie
		for indexRange := range gopart.Partition(len(players), 10) {
			go save(players[indexRange.Low:indexRange.High], t)
		}
		c <- true
	}()

	<-c

	return nil
}

func ratePlayers(config model.ScoreConfig, t *trie.Trie) []model.Player {
	return scorePlayers(config, t)
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
