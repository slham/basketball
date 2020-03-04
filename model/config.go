package model

import (
	"math"
	"time"
)

type ScoreConfig struct {
	Min float32 `yaml:"min"  validate:"min=0,max=10"`
	Fgm float32 `yaml:"fgm"  validate:"min=0,max=10"`
	Fga float32 `yaml:"fga"  validate:"min=0,max=10"`
	Fgp float32 `yaml:"fgp"  validate:"min=0,max=10"`
	Ftm float32 `yaml:"ftm"  validate:"min=0,max=10"`
	Fta float32 `yaml:"fta"  validate:"min=0,max=10"`
	Ftp float32 `yaml:"ftp"  validate:"min=0,max=10"`
	Tpm float32 `yaml:"tpm"  validate:"min=0,max=10"`
	Tpa float32 `yaml:"tpa"  validate:"min=0,max=10"`
	Tpp float32 `yaml:"tpp"  validate:"min=0,max=10"`
	Reb float32 `yaml:"reb"  validate:"min=0,max=10"`
	Ass float32 `yaml:"ass"  validate:"min=0,max=10"`
	Stl float32 `yaml:"stl"  validate:"min=0,max=10"`
	Blk float32 `yaml:"blk"  validate:"min=0,max=10"`
	Tvs float32 `yaml:"tvs"  validate:"min=0,max=10"`
	Dds float32 `yaml:"dds"  validate:"min=0,max=10"`
	Pts float32 `yaml:"pts"  validate:"min=0,max=10"`
}

func (config *ScoreConfig) Score(player *Player) {
	score := float32(0.0)
	numGames := float32(player.Gms)

	score += float32(player.Min) / numGames * config.Min
	score += player.Fgm / numGames * config.Fgm
	score += player.Fga / numGames * config.Fga
	score += player.Fgp * config.Fgp
	score += player.Ftm / numGames * config.Ftm
	score += player.Fta / numGames * config.Fta
	score += player.Ftp * config.Ftp
	score += player.Tpm / numGames * config.Ftm
	score += player.Tpa / numGames * config.Tpa
	score += player.Tpp * config.Tpp
	score += player.Reb * config.Reb
	score += player.Ass * config.Ass
	score += player.Stl * config.Stl
	score += player.Blk * config.Blk
	score += player.Tvs * config.Tvs
	score += player.Dds / numGames * config.Dds
	score += player.Pts / numGames * config.Pts

	player.Score = round(score, 2)
	player.UpdatedDateTime = time.Now()
}

func round(val float32, precision int) float32 {
	pow10 := math.Pow10(precision)
	return float32(math.Round(float64(val)*pow10) / pow10)
}
