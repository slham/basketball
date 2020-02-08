package app

import (
	"basketball/model"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestValidateScoreConfig(t *testing.T) {
	tables := []struct {
		config  model.ScoreConfig
		message string
	}{
		{model.ScoreConfig{
			Min: 1.0,
			Fgm: 1.0,
			Fga: 1.0,
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
		},
			"Key: 'ScoreConfig.Fgp' Error:Field validation for 'Fgp' failed on the 'required' tag",
		},
		{model.ScoreConfig{
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
		},
			"",
		},
	}

	for _, table := range tables {
		err := validateScoreConfig(table.config)
		if err != nil {
			assert.Equal(t, table.message, err.Error())
		}
	}
}
