package storage

import (
	"basketball/model"
	"github.com/golang-collections/collections/trie"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	player := model.Player{
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
	bytes, err := hash(player)
	if err != nil {
		t.Error(err)
	}

	for i, b := range []byte{uint8(224), uint8(152), uint8(110), uint8(144), uint8(160), uint8(90), uint8(177), uint8(38), uint8(81), uint8(74), uint8(18), uint8(31), uint8(14), uint8(197), uint8(131), uint8(0)} {
		assert.Equal(t, b, bytes[i])
	}
}

func TestScorePlayers(t *testing.T) {
	players := []model.Player{
		model.Player{
			Id:  0,
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
		},
		model.Player{
			Id:  1,
			Gms: 32,
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
			Dds: 3.0,
			Pts: 172.0,
		},
		model.Player{
			Id:  2,
			Gms: 12,
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
			Dds: 3.0,
			Pts: 172.0,
		},
	}
	testTrie := trie.New()
	testTrie.Init()
	for _, player := range players {

		now := time.Now()
		player.CreatedDateTime = now
		player.UpdatedDateTime = now
		key, err := hash(player)
		if err != nil {
			t.Error(err)
		}
		testTrie.Insert(key, player)
	}

	config := model.ScoreConfig{
		Min: 1.0,
		Fgm: 1.0,
		Fga: 1.0,
		Fgp: 6.0,
		Ftm: 1.0,
		Fta: 1.0,
		Ftp: 1.0,
		Tpm: 1.0,
		Tpa: 1.0,
		Tpp: 1.0,
		Reb: 1.0,
		Ass: 1.0,
		Stl: 3.0,
		Blk: 1.0,
		Tvs: 3.0,
		Dds: 1.0,
		Pts: 1.0,
	}

	res := ScorePlayers(config, testTrie)
	assert.Equal(t, 2, len(res))
	p1, p2 := res[0], res[1]
	assert.Equal(t, true, p1.Id == 1)
	assert.Equal(t, true, p2.Id == 0)
	assert.Equal(t, true, p1.Score > 0)
	assert.Equal(t, true, p2.Score > 0)
	assert.Equal(t, true, p1.Score > p2.Score)
}

func TestFillTeam(t *testing.T) {
	players := []model.Player{
		{Name: "James Harden", Position: "PG"},
		{Name: "Damian Lillard", Position: "PG"},
		{Name: "Trae", Position: "PG"},
		{Name: "Kyrie Irving", Position: "PG"},
		{Name: "Luka Doncic", Position: "PG"},
		{Name: "Bradley Beal", Position: "SG"},
		{Name: "Kawhi Leonard", Position: "SF"},
		{Name: "Karl-Anthony Towns", Position: "C"},
		{Name: "LeBron James", Position: "SF"},
		{Name: "Giannis Antetokounmpo", Position: "PF"},
		{Name: "Anthony Davis", Position: "PF"},
		{Name: "Russell Westbrook", Position: "PG"},
		{Name: "Brandon Ingram", Position: "SF"},
		{Name: "Nikola Jokic", Position: "C"},
	}

	team := FillTeam(nil, players)
	assert.Equal(t, 10, len(team))
}
