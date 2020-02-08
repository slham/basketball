package app

import (
	"basketball/model"
	"gopkg.in/go-playground/validator.v9"
)

func validateScoreConfig(scoreConfig model.ScoreConfig) error {
	v := validator.New()
	err := v.Struct(scoreConfig)
	if err != nil {
		return err
	}

	return nil
}
