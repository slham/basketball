package model

import (
	"math"
	"time"
)

type ScoreConfig struct {
	Min float32 `yaml:"min"  validate:"min=0,max=10,required"`
	Fgm float32 `yaml:"fgm"  validate:"min=0,max=10,required"`
	Fga float32 `yaml:"fga"  validate:"min=0,max=10,required"`
	Fgp float32 `yaml:"fgp"  validate:"min=0,max=10,required"`
	Ftm float32 `yaml:"ftm"  validate:"min=0,max=10,required"`
	Fta float32 `yaml:"fta"  validate:"min=0,max=10,required"`
	Ftp float32 `yaml:"ftp"  validate:"min=0,max=10,required"`
	Tpm float32 `yaml:"tpm"  validate:"min=0,max=10,required"`
	Tpa float32 `yaml:"tpa"  validate:"min=0,max=10,required"`
	Tpp float32 `yaml:"tpp"  validate:"min=0,max=10,required"`
	Reb float32 `yaml:"reb"  validate:"min=0,max=10,required"`
	Ass float32 `yaml:"ass"  validate:"min=0,max=10,required"`
	Stl float32 `yaml:"stl"  validate:"min=0,max=10,required"`
	Blk float32 `yaml:"blk"  validate:"min=0,max=10,required"`
	Tvs float32 `yaml:"tvs"  validate:"min=0,max=10,required"`
	Dds float32 `yaml:"dds"  validate:"min=0,max=10,required"`
	Pts float32 `yaml:"pts"  validate:"min=0,max=10,required"`
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
