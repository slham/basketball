package model

import (
	"time"
)

type Player struct {
	Id              int       `json:"PlayerID" yaml:"id"`
	TeamId          int       `json:"TeamID" yaml:"teamId"`
	Name            string    `json:"Name" yaml:"name"`
	Position        string    `json:"Position" yaml:"pos"`
	Min             int       `json:"Minutes" yaml:"min"`
	Fgm             float32   `json:"FieldGoalsMade" yaml:"fgm"`
	Fga             float32   `json:"FieldGoalsAttempted" yaml:"fga"`
	Fgp             float32   `json:"FieldGoalsPercentage" yaml:"fgp"`
	Ftm             float32   `json:"FreeThrowsMade" yaml:"ftm"`
	Fta             float32   `json:"FreeThrowsAttempted" yaml:"fta"`
	Ftp             float32   `json:"FreeThrowsPercentage" yaml:"ftp"`
	Tpm             float32   `json:"ThreePointersMade" yaml:"tpm"`
	Tpa             float32   `json:"ThreePointersAttempted" yaml:"tpa"`
	Tpp             float32   `json:"ThreePointersPercentage" yaml:"tpp"`
	Reb             float32   `json:"TotalReboundsPercentage" yaml:"reb"`
	Ass             float32   `json:"AssistsPercentage" yaml:"ass"`
	Stl             float32   `json:"StealsPercentage" yaml:"stl"`
	Blk             float32   `json:"BlocksPercentage" yaml:"bks"`
	Tvs             float32   `json:"TurnOversPercentage" yaml:"tvs"`
	Dds             float32   `json:"DoubleDoubles" yaml:"dds"`
	Pts             float32   `json:"Points" yaml:"pts"`
	Gms             int       `json:"Games" yaml:"gms"`
	CreatedDateTime time.Time `yaml:"createdDateTime"`
	UpdatedDateTime time.Time `yaml:"updatedDateTime"`
	Score           float32   `yaml:"score"`
}
