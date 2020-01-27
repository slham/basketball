package main

import (
	"basketball/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"gopkg.in/yaml.v2"
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
	var store = make(map[int]float32, 0)
	for _, player := range players {
		scoreConfig.Score(&player)
		store[player.Id] = player.Score
	}

	log.Println(fmt.Sprintf("store: %v", store))

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
