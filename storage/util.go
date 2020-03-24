package storage

import (
	"basketball/model"
	"context"
	"crypto/md5"
	"github.com/golang-collections/collections/trie"
	"github.com/meirf/gopart"
	"github.com/slham/toolbelt/coll"
	"github.com/slham/toolbelt/l"
	"gopkg.in/yaml.v2"
	"sort"
)

type ByScore []model.Player

func (score ByScore) Len() int           { return len(score) }
func (score ByScore) Swap(i, j int)      { score[i], score[j] = score[j], score[i] }
func (score ByScore) Less(i, j int) bool { return score[i].Score > score[j].Score }

type Position struct {
	Name             string
	AllowedPositions []string
	Filled           bool
}

func hash(it interface{}) ([16]byte, error) {
	bytes, err := yaml.Marshal(it)
	if err != nil {
		return [16]byte{}, err
	}
	return md5.Sum(bytes), nil
}

func ScorePlayers(config model.ScoreConfig, t *trie.Trie) []model.Player {
	var players = make([]model.Player, 0)
	if t == nil || t.Len() == 0 {
		return players
	}

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

func FillTeam(ctx context.Context, players []model.Player) []model.Player {
	positions := []Position{
		{Name: "PG", AllowedPositions: []string{"PG"}, Filled: false},
		{Name: "SG", AllowedPositions: []string{"SG"}},
		{Name: "SF", AllowedPositions: []string{"SF"}},
		{Name: "PF", AllowedPositions: []string{"PF"}},
		{Name: "C", AllowedPositions: []string{"C"}},
		{Name: "G", AllowedPositions: []string{"PG", "SG"}},
		{Name: "F", AllowedPositions: []string{"SF", "PF"}},
		{Name: "UTIL", AllowedPositions: []string{"PG", "SG", "SF", "PF", "C"}},
		{Name: "BENCH", AllowedPositions: []string{"PG", "SG", "SF", "PF", "C"}},
		{Name: "BENCH", AllowedPositions: []string{"PG", "SG", "SF", "PF", "C"}},
		{Name: "BENCH", AllowedPositions: []string{"PG", "SG", "SF", "PF", "C"}},
	}
	team := make([]model.Player, 0)
	positionCaps := map[string]int8{
		"PG": int8(3),
		"SG": int8(3),
		"SF": int8(3),
		"PF": int8(3),
		"C":  int8(2),
	}
	positionCount := map[string]int8{
		"PG": int8(0),
		"SG": int8(0),
		"SF": int8(0),
		"PF": int8(0),
		"C":  int8(0),
	}

	for indexRange := range gopart.Partition(len(players), 50) {
		if len(team) == 10 {
			l.Debug(ctx, "team is full")
			break
		}
		for _, player := range players[indexRange.Low:indexRange.High] {
			if len(team) == 10 {
				l.Debug(ctx, "team is full")
				break
			}
			for i := 0; i < len(positions); i++ {
				position := positions[i]
				if positionCount[player.Position] == positionCaps[player.Position] {
					continue
				}
				if coll.Include(position.AllowedPositions, player.Position) {
					team = append(team, player)
					positionCount[player.Position]++
					positions = append(positions[:i], positions[i+1:]...)
					l.Debug(ctx, "putting %s at %s", player.Name, position.Name)
					break
				}
			}
		}
	}

	return team
}
