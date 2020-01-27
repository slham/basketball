package model

type ScoreConfig struct {
	Min             int       `yaml:"min"`
	Fgm             float32   `yaml:"fgm"`
	Fga             float32   `yaml:"fga"`
	Fgp             float32   `yaml:"fgp"`
	Ftm             float32   `yaml:"ftm"`
	Fta             float32   `yaml:"fta"`
	Ftp             float32   `yaml:"ftp"`
	Tpm             float32   `yaml:"tpm"`
	Tpa             float32   `yaml:"tpa"`
	Tpp             float32   `yaml:"tpp"`
	Reb             float32   `yaml:"reb"`
	Ass             float32   `yaml:"ass"`
	Stl             float32   `yaml:"stl"`
	Blk             float32   `yaml:"bks"`
	Turnovers       float32   `yaml:"tvs"`
	Dds             float32   `yaml:"dds"`
	Pts             float32   `yaml:"pts"`
}
