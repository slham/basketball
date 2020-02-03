package model

import (
	"time"
)

type Player struct {
	Id              int       `json:"PlayerID" yaml:"id"`
	TeamId          int       `json:"TeamID" yaml:"teamId,omitempty"`
	Name            string    `json:"Name" yaml:"name"`
	Position        string    `json:"Position" yaml:"pos,omitempty"`
	Min             int       `json:"Minutes" yaml:"min,omitempty"`
	Fgm             float32   `json:"FieldGoalsMade" yaml:"fgm,omitempty"`
	Fga             float32   `json:"FieldGoalsAttempted" yaml:"fga,omitempty"`
	Fgp             float32   `json:"FieldGoalsPercentage" yaml:"fgp,omitempty"`
	Ftm             float32   `json:"FreeThrowsMade" yaml:"ftm,omitempty"`
	Fta             float32   `json:"FreeThrowsAttempted" yaml:"fta,omitempty"`
	Ftp             float32   `json:"FreeThrowsPercentage" yaml:"ftp,omitempty"`
	Tpm             float32   `json:"ThreePointersMade" yaml:"tpm,omitempty"`
	Tpa             float32   `json:"ThreePointersAttempted" yaml:"tpa,omitempty"`
	Tpp             float32   `json:"ThreePointersPercentage" yaml:"tpp,omitempty"`
	Reb             float32   `json:"TotalReboundsPercentage" yaml:"reb,omitempty"`
	Ass             float32   `json:"AssistsPercentage" yaml:"ass,omitempty"`
	Stl             float32   `json:"StealsPercentage" yaml:"stl,omitempty"`
	Blk             float32   `json:"BlocksPercentage" yaml:"bks,omitempty"`
	Tvs             float32   `json:"TurnOversPercentage" yaml:"tvs,omitempty"`
	Dds             float32   `json:"DoubleDoubles" yaml:"dds,omitempty"`
	Pts             float32   `json:"Points" yaml:"pts,omitempty"`
	Gms             int       `json:"Games" yaml:"gms,omitempty"`
	CreatedDateTime time.Time `yaml:"createdDateTime"`
	UpdatedDateTime time.Time `yaml:"updatedDateTime"`
	Score           float32   `yaml:"score"`
}
