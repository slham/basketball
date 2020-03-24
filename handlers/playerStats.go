package handlers

import (
	"basketball/model"
	"basketball/storage"
	"basketball/valid"
	"github.com/slham/toolbelt/l"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type RatePlayersResponse struct {
	Team    []model.Player
	Players []model.Player
}

func RatePlayers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l.Debug(ctx, "ratings request received")

	var scoreConfig model.ScoreConfig
	err := yaml.NewDecoder(r.Body).Decode(&scoreConfig)
	if err != nil {
		l.Error(ctx, "reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid payload"))
		return
	}

	err = valid.ValidateScoreConfig(scoreConfig)
	if err != nil {
		l.Error(ctx, "invalid score config %v: %v", scoreConfig, err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		return
	}

	l.Debug(ctx, "rating players for current config: %v", scoreConfig)

	//rate players using config
	players := storage.ScorePlayers(scoreConfig, storage.Store())

	//make suggested team
	team := storage.FillTeam(ctx, players)

	response := RatePlayersResponse{
		Team:    team,
		Players: players,
	}

	//marshall response body
	bytes, err := yaml.Marshal(response)
	if err != nil {
		l.Error(ctx, "responding with rated players: %v", err)
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
	l.Info(ctx, "updating player stats data")

	playerBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Error(ctx, "reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("unable to read payload"))
		return
	}

	err = storage.UnmarshalAndSavePlayers(ctx, playerBytes, storage.Store())
	if err != nil {
		l.Error(ctx, "updating player stats data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("failed store players"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write([]byte("successfully stored players"))
	}
}
