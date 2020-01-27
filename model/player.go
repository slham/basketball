package model

import (
	"time"
)

type Player struct {
	Id              int    `json:"PlayerID" yaml:"id"`
	TeamId          int    `json:"TeamID" yaml:"teamId"`
	Name            string    `json:"Name" yaml:"name"`
	Position        string    `json:"Position" yaml:"pos"`
	Min             int       `json:"Minutes" yaml:"min"`
	Sec             int       `json:"Seconds"`
	Fgm             float32   `json:"FieldGoalsMade" yaml:"fgm"`
	Fga             float32   `json:"FieldGoalsAttempted" yaml:"fga"`
	Fgp             float32   `json:"FieldGoalsPercentage" yaml:"fgp"`
	Ftm             float32   `json:"FreeThrowsMade" yaml:"ftm"`
	Fta             float32   `json:"FreeThrowsAttempted" yaml:"fta"`
	Ftp             float32   `json:"FreeThrowsPercentage" yaml:"ftp"`
	Tpm             float32   `json:"ThreePointersMade" yaml:"tpm"`
	Tpa             float32   `json:"ThreePointersAttempted" yaml:"tpa"`
	Tpp             float32   `json:"ThreePointersPercentage" yaml:"tpp"`
	Reb             float32   `json:"Rebounds" yaml:"reb"`
	Ass             float32   `json:"Assists" yaml:"ass"`
	Stl             float32   `json:"Steals" yaml:"stl"`
	Blk             float32   `json:"Blocks" yaml:"blocks"`
	Turnovers       float32   `json:"Turnovers" yaml:"turnovers"`
	Dds             float32   `yaml:"DoubleDoubles"`
	Pts             float32   `json:"Points" yaml:"pts"`
	CreatedDateTime time.Time `yaml:"createdDateTime"`
	UpdatedDateTime time.Time `yaml:"updatedDateTime"`
}
