package model

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestScoreConfig_Score(t *testing.T) {
	player := Player{
		Gms: 42,
		Min: 621,
		Fgm: 68.3,
		Fga: 132.1,
		Fgp: 58.9,
		Ftm: 35.3,
		Fta: 53.5,
		Ftp: 75.2,
		Tpm: 0.0,
		Tpa: 3.9,
		Tpp: 0.0,
		Reb: 140.1,
		Ass: 29.6,
		Stl: 20.5,
		Blk: 17.1,
		Tvs: 12.5,
		Dds: 0.0,
		Pts: 172.0,
	}
	config := ScoreConfig{
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
	config.Score(&player)
	assert.Equal(t, float32(379.75955), player.Score)
}
