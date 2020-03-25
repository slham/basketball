package handlers

import (
	"basketball/env"
	"basketball/model"
	"basketball/storage"
	"bytes"
	"github.com/slham/toolbelt/l"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRatePlayersNoPlayers(t *testing.T) {
	scoreConfig := model.ScoreConfig{
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
	configBytes, err := yaml.Marshal(scoreConfig)
	if err != nil {
		t.Error(err)
		return
	}

	req, err := http.NewRequest("POST", "/ratings", bytes.NewReader(configBytes))
	if handelError(err, t) {
		return
	}

	initializeContext(t, false)

	rr := httptest.NewRecorder()
	withLogging := l.Logging(http.HandlerFunc(RatePlayers))
	withLogging.ServeHTTP(rr, req)

	var ratingsRes RatePlayersResponse
	err = yaml.Unmarshal(rr.Body.Bytes(), &ratingsRes)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 0, len(ratingsRes.Players))
}

func handelError(err error, t *testing.T) bool {
	if err != nil {
		t.Error(err)
		return true
	}
	return false
}

func TestRatePlayers(t *testing.T) {
	scoreConfig := model.ScoreConfig{
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
	configBytes, err := yaml.Marshal(scoreConfig)
	if err != nil {
		t.Error(err)
		return
	}

	req, err := http.NewRequest("POST", "/ratings", bytes.NewReader(configBytes))
	if err != nil {
		t.Error(err)
		return
	}

	initializeContext(t, true)

	rr := httptest.NewRecorder()
	withLogging := l.Logging(http.HandlerFunc(RatePlayers))
	withLogging.ServeHTTP(rr, req)

	var ratingsRes RatePlayersResponse
	err = yaml.Unmarshal(rr.Body.Bytes(), &ratingsRes)
	if err != nil {
		t.Error(err)
		return
	}

	players := ratingsRes.Players
	team := ratingsRes.Team

	assert.Equal(t, http.StatusOK, rr.Code)
	for i := 0; i < len(players)-1; i++ {
		assert.Equal(t, players[i].Score > players[i+1].Score, true)
	}

	for i := 0; i < len(team)-1; i++ {
		assert.Equal(t, team[i].Score > team[i+1].Score, true)
	}
}

func initializeContext(t *testing.T, withPlayers bool) {
	l.Initialize(l.DEBUG)
	if withPlayers {
		var config env.Config
		config.Storage.FileName = "1583510437-test.yaml"
		config.Env = "DEV"
		ok := storage.Initialize(config)
		if !ok {
			t.Error("could not set up storage")
		}
	}
	return
}

func TestRatePlayersInvalidPayload(t *testing.T) {
	req, err := http.NewRequest("POST", "/ratings", bytes.NewReader([]byte("failure")))
	if handelError(err, t) {
		return
	}

	initializeContext(t, true)

	rr := httptest.NewRecorder()
	withLogging := l.Logging(http.HandlerFunc(RatePlayers))
	withLogging.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "invalid payload", rr.Body.String())
}

func TestRatePlayersInvalidScoreConfig(t *testing.T) {
	scoreConfig := model.ScoreConfig{
		Min: -1.0,
		Fgm: -1.0,
		Fga: -1.0,
		Fgp: -1.0,
		Ftm: -1.0,
		Fta: -1.0,
		Ftp: 1.0,
		Tpm: -1.0,
		Tpa: 1.0,
		Tpp: 19.0,
		Reb: 16.0,
		Ass: 1.0,
		Stl: 1.0,
		Blk: 1.0,
		Tvs: 1.0,
		Dds: 1.0,
		Pts: 1.0,
	}
	configBytes, err := yaml.Marshal(scoreConfig)
	if err != nil {
		t.Error(err)
		return
	}

	req, err := http.NewRequest("POST", "/ratings", bytes.NewReader(configBytes))
	if err != nil {
		t.Error(err)
		return
	}

	initializeContext(t, true)

	rr := httptest.NewRecorder()
	withLogging := l.Logging(http.HandlerFunc(RatePlayers))
	withLogging.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "all fields must be populated with a number between 0.0 and 10.0", rr.Body.String())
}

func TestStorePlayers(t *testing.T) {
	players := []model.Player{
		{Id: 0}, {Id: 1}, {Id: 2},
	}
	playerBytes, err := yaml.Marshal(players)
	if handelError(err, t) {
		return
	}

	req, err := http.NewRequest("PUT", "/players", bytes.NewReader(playerBytes))
	if handelError(err, t) {
		return
	}

	initializeContext(t, true)
	rr := httptest.NewRecorder()
	withLogging := l.Logging(http.HandlerFunc(StorePlayers))
	withLogging.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "successfully stored players", rr.Body.String())
}
