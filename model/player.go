package model

import (
	"time"
)

type Player struct {
	Id              int       `json:"PlayerID" yaml:"id"`
	TeamId          int       `json:"TeamID,omitempty" yaml:"teamId,omitempty"`
	Name            string    `json:"Name" yaml:"name"`
	Position        string    `json:"Position,omitempty" yaml:"pos,omitempty"`
	Min             int       `json:"Minutes,omitempty" yaml:"min,omitempty"`
	Fgm             float32   `json:"FieldGoalsMade,omitempty" yaml:"fgm,omitempty"`
	Fga             float32   `json:"FieldGoalsAttempted,omitempty" yaml:"fga,omitempty"`
	Fgp             float32   `json:"FieldGoalsPercentage,omitempty" yaml:"fgp,omitempty"`
	Ftm             float32   `json:"FreeThrowsMade,omitempty" yaml:"ftm,omitempty"`
	Fta             float32   `json:"FreeThrowsAttempted,omitempty" yaml:"fta,omitempty"`
	Ftp             float32   `json:"FreeThrowsPercentage,omitempty" yaml:"ftp,omitempty"`
	Tpm             float32   `json:"ThreePointersMade,omitempty" yaml:"tpm,omitempty"`
	Tpa             float32   `json:"ThreePointersAttempted,omitempty" yaml:"tpa,omitempty"`
	Tpp             float32   `json:"ThreePointersPercentage,omitempty" yaml:"tpp,omitempty"`
	Reb             float32   `json:"TotalReboundsPercentage,omitempty" yaml:"reb,omitempty"`
	Ass             float32   `json:"AssistsPercentage,omitempty" yaml:"ass,omitempty"`
	Stl             float32   `json:"StealsPercentage,omitempty" yaml:"stl,omitempty"`
	Blk             float32   `json:"BlocksPercentage,omitempty" yaml:"bks,omitempty"`
	Tvs             float32   `json:"TurnOversPercentage,omitempty" yaml:"tvs,omitempty"`
	Dds             float32   `json:"DoubleDoubles,omitempty" yaml:"dds,omitempty"`
	Pts             float32   `json:"Points,omitempty" yaml:"pts,omitempty"`
	Gms             int       `json:"Games,omitempty" yaml:"gms,omitempty"`
	CreatedDateTime time.Time `yaml:"createdDateTime"`
	UpdatedDateTime time.Time `yaml:"updatedDateTime"`
	Score           float32   `yaml:"score"`
}