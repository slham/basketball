package model

import "time"

type ScoreConfig struct {
	Min float32 `yaml:"min"`
	Fgm float32 `yaml:"fgm"`
	Fga float32 `yaml:"fga"`
	Fgp float32 `yaml:"fgp"`
	Ftm float32 `yaml:"ftm"`
	Fta float32 `yaml:"fta"`
	Ftp float32 `yaml:"ftp"`
	Tpm float32 `yaml:"tpm"`
	Tpa float32 `yaml:"tpa"`
	Tpp float32 `yaml:"tpp"`
	Reb float32 `yaml:"reb"`
	Ass float32 `yaml:"ass"`
	Stl float32 `yaml:"stl"`
	Blk float32 `yaml:"bks"`
	Tvs float32 `yaml:"tvs"`
	Dds float32 `yaml:"dds"`
	Pts float32 `yaml:"pts"`
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

	player.Score = score
	player.UpdatedDateTime = time.Now()
}
