package valid

import (
	"basketball/model"
	"gopkg.in/go-playground/assert.v1"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestValidateScoreConfig(t *testing.T) {
	tables := []struct {
		config  model.ScoreConfig
		reason string
	}{
		{model.ScoreConfig{
			Min: -1.0,
			Fgm: -1.0,
			Fga: -1.0,
			Fgp: -1.0,
			Ftm: -1.0,
			Fta: -1.0,
			Ftp: -1.0,
			Tpm: -1.0,
			Tpa: -1.0,
			Tpp: -1.0,
			Reb: -1.0,
			Ass: -1.0,
			Stl: -1.0,
			Blk: -1.0,
			Tvs: -1.0,
			Dds: -1.0,
			Pts: -1.0,
		},
		"min",
		},
		{model.ScoreConfig{
			Min: 18.0,
			Fgm: 16.0,
			Fga: 13.0,
			Fgp: 15.0,
			Ftm: 17.0,
			Fta: 19.0,
			Ftp: 10.0,
			Tpm: 12.0,
			Tpa: 14.0,
			Tpp: 11.0,
			Reb: 15.0,
			Ass: 17.0,
			Stl: 13.0,
			Blk: 15.0,
			Tvs: 18.0,
			Dds: 19.0,
			Pts: 100.0,
		},
			"max",
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
		err := ValidateScoreConfig(table.config)
		if err != nil {
			for _, tag := range err.(validator.ValidationErrors) {
				assert.Equal(t, table.reason, validator.FieldError(tag).Tag())
			}
		}else {
			t.Log("valid ScoreConfig")
		}
	}
}
