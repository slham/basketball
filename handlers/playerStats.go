package handlers

import (
	"basketball/model"
	"basketball/storage"
	"basketball/valid"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

func RatePlayers(w http.ResponseWriter, r *http.Request) {
	log.Println("ratings request received")

	var scoreConfig model.ScoreConfig
	err := yaml.NewDecoder(r.Body).Decode(&scoreConfig)
	if err != nil {
		msg := fmt.Sprintf("error reading request body: %v", err)
		log.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	err = valid.ValidateScoreConfig(scoreConfig)
	if err != nil {
		msg := fmt.Sprintf("invalid score config %v: %v", scoreConfig, err)
		log.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	log.Println(fmt.Sprintf("rating players for current config: %v", scoreConfig))
	//rate players using config
	players := storage.ScorePlayers(scoreConfig, storage.Store())

	//marshall response body
	bytes, err := yaml.Marshal(players)
	if err != nil {
		msg := fmt.Sprintf("error responding with rated players: %v", err)
		log.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed to rate players"))
		return
	}

	//respond with players
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/yaml")
	_, _ = w.Write(bytes)
}

func StorePlayers(w http.ResponseWriter, r *http.Request) {
	log.Println("updating player stats data")

	playerBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("error reading request body: %v", err)
		log.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	err = storage.UnmarshalAndSavePlayers(playerBytes, storage.Store())
	if err != nil {
		log.Println("error updating player stats data")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed store players"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write([]byte("successfully stored players"))
	}
}
