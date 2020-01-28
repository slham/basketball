package main

import (
	"basketball/model"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/meirf/gopart"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"github.com/golang-collections/collections/trie"
	"time"
)

func main() {
	var players = make([]model.Player, 0)
	scoreConfig := loadScoreConfig()

	jsonFile, err := os.Open("sample.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened sample.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &players)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Unmarshalled sample.json")
	//var store = make(map[int]float32, 0)
	myTrie := trie.New()
	myTrie.Init()
	for indexRange := range gopart.Partition(len(players), 10){
		go scorePlayers(scoreConfig, players[indexRange.Low:indexRange.High], myTrie)
	}


	<-time.After(time.Second * 5)
	log.Println(fmt.Sprintf("store: %v", myTrie.String()))

}

func scorePlayers(scoreConfig model.ScoreConfig, players []model.Player, t *trie.Trie){
	for _, player := range players {
		scoreConfig.Score(&player)
		key := hash(player)
		storePlayer(player)
		t.Insert(key, player)
	}
}

func storePlayer(player model.Player) {
	file, err := os.Create(fmt.Sprintf("data/%d_%v.yaml", player.Id, player.UpdatedDateTime.Unix()))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := yaml.Marshal(player)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(bytes)
	file.Sync()
}

func hash(player model.Player) [16]byte {
	bytes, err := yaml.Marshal(player)
	if err != nil {
		log.Fatal(err)
	}
	return md5.Sum(bytes)
}

func loadScoreConfig() model.ScoreConfig {
	//scoreConfigYaml, err := os.Open("score-config.yaml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer scoreConfigYaml.Close()
	//
	//bytes, err := ioutil.ReadAll(scoreConfigYaml)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = yaml.Unmar
	return model.ScoreConfig{
		Min: 1.0,
		Fgm: 1.0,
		Fga: 1.0,
		Fgp: 1.0,
		Ftm: 1.0,
		Fta: 1.0,
		Ftp: 1.0,
		Tpm: 1.0,
		Tpa: 1.0,
		Tpp: 1.0,
		Reb: 1.0,
		Ass: 1.0,
		Stl: 1.0,
		Blk: 1.0,
		Tvs: 1.0,
		Dds: 1.0,
		Pts: 1.0,
	}
}
