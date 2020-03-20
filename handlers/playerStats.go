package handlers

import (
	"basketball/model"
	"basketball/storage"
	"basketball/valid"
	"github.com/slham/toolbelt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func RatePlayers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	toolbelt.Debug(ctx, "ratings request received")

	var scoreConfig model.ScoreConfig
	err := yaml.NewDecoder(r.Body).Decode(&scoreConfig)
	if err != nil {
		toolbelt.Error(ctx, "reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	err = valid.ValidateScoreConfig(scoreConfig)
	if err != nil {
		toolbelt.Error(ctx, "invalid score config %v: %v", scoreConfig, err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	toolbelt.Debug(ctx, "rating players for current config: %v", scoreConfig)

	//rate players using config
	players := storage.ScorePlayers(scoreConfig, storage.Store())

	//marshall response body
	bytes, err := yaml.Marshal(players)
	if err != nil {
		toolbelt.Error(ctx, "responding with rated players: %v", err)
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
	ctx := r.Context()
	toolbelt.Info(ctx, "updating player stats data")

	playerBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		toolbelt.Error(ctx, "reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	err = storage.UnmarshalAndSavePlayers(ctx, playerBytes, storage.Store())
	if err != nil {
		toolbelt.Error(ctx, "updating player stats data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed store players"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write([]byte("successfully stored players"))
	}
}
