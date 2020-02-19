package app

import (
	"basketball/model"
	"encoding/json"
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
func fetchFromSource(t *trie.Trie) error {
	url := fmt.Sprintf("https://api.sportsdata.io/v3/nba/stats/json/PlayerSeasonStats/2020?key=%s", os.Getenv("NBA_API_KEY"))
	log.Println(fmt.Sprintf("sending to %v", url))
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	players := make([]model.Player, 0)

	err = json.Unmarshal(body, &players)
	if err != nil {
		return err
	}

	//hash and store in trie
	for indexRange := range gopart.Partition(len(players), 10) {
		go save(players[indexRange.Low:indexRange.High], t)
	}


	return nil
}

func ratePlayers(config model.ScoreConfig, t *trie.Trie) []model.Player {
	return scorePlayers(config, t)
}

func save(players []model.Player, t *trie.Trie) {
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
