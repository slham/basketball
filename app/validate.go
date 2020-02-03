package app

import (
	"basketball/model"
	"gopkg.in/yaml.v2"
	"net/http"
	//"gopkg.in/go-playground/validator.v9"
)

func validateScoreConfig(r *http.Request) (model.ScoreConfig, error){
	var scoreConfig model.ScoreConfig
	err := yaml.NewDecoder(r.Body).Decode(&scoreConfig)
	if err != nil {
		return scoreConfig, err
	}
	//validate scoreConfig
	return scoreConfig, nil
}
