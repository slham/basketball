package app

import (
	"basketball/model"
	"crypto/md5"
	"github.com/golang-collections/collections/trie"
	"gopkg.in/yaml.v2"
	"log"
	"sort"
)

type ByScore []model.Player

func (score ByScore) Len() int           { return len(score) }
func (score ByScore) Swap(i, j int)      { score[i], score[j] = score[j], score[i] }
func (score ByScore) Less(i, j int) bool { return score[i].Score > score[j].Score }

func hash(player model.Player) ([16]byte, error) {
	bytes, err := yaml.Marshal(player)
	if err != nil {
		log.Println("unable to hash player")
		return [16]byte{}, err
	}
	return md5.Sum(bytes), nil
}

func scorePlayers(config model.ScoreConfig, t *trie.Trie) []model.Player {
	var players = make([]model.Player, 0)

	t.Do(func(k, v interface{}) bool {
		player := v.(model.Player)
		if player.Gms >= 15 {
			config.Score(&player)
			players = append(players, model.Player{
				Id:              player.Id,
				Name:            player.Name,
				Position:        player.Position,
				Score:           player.Score,
				CreatedDateTime: player.CreatedDateTime,
				UpdatedDateTime: player.UpdatedDateTime,
			})
		}
		return true
	})

	sort.Sort(ByScore(players))

	return players
}
