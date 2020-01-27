package model

import (
	"time"
)

type Player struct {
	Id              string       `json:"playerId" yaml:"id"`
	TeamId          string       `json:"teamId" yaml:"teamId"`
	NumGames        int       `yaml:"numGames"`
	FirstName       string    `json:"firstName" yaml:"firstName"`
	LastName        string    `json:"lastName" yaml:"firstName"`
	Position        string    `json:"pos" yaml:"pos"`
	Min             string    `json:"min" yaml:"min"`
	Fgm             int       `json:"fgm" yaml:"fgm"`
	Fga             int       `json:"fga" yaml:"fga"`
	Fgp             float32   `json:"fgp" yaml:"fgp"`
	Ftm             int       `json:"ftm" yaml:"ftm"`
	Fta             int       `json:"fta" yaml:"fta"`
	Ftp             float32   `json:"ftp" yaml:"ftp"`
	Tpm             int       `json:"tpm" yaml:"tpm"`
	Tpa             int       `json:"tpa" yaml:"tpa"`
	Tpp             float32   `json:"tpp" yaml:"tpp"`
	Reb             int       `json:"totReb" yaml:"reb"`
	Ass             int       `json:"assists" yaml:"ass"`
	Stl             int       `json:"steals" yaml:"stl"`
	Blk             int       `json:"blocks" yaml:"blocks"`
	Turnovers       int       `json:"turnovers" yaml:"turnovers"`
	Dds             int       `yaml:"dds"`
	Pts             int       `json:"points" yaml:"pts"`
	CreatedDateTime time.Time `json:"createdDateTime" yaml:"createdDateTime"`
	UpdatedDateTime time.Time `json:"updatedDateTime" yaml:"updatedDateTime"`
}

func (p Player) CheckDds(){
	check := 0
	checkPass := func(p Player) {
		if check >1 {
			p.Dds++
		}

	}
	if p.Pts >= 10{
		check++
	}
	if p.Reb >= 10{
		check++
		checkPass(p)
	}
	if p.Ass >= 10{
		check++
		checkPass(p)
	}
	if p.Stl >= 10{
		check++
		checkPass(p)
	}
	if p.Blk >= 10{
		check++
		checkPass(p)
	}
}

//func Merge(a, b Player) (c Player) {
//	if a.Id == -1 || b.Id == -1 || a.Id != b.Id ||
//		a.CreatedDateTime == (time.Time{}) || a.UpdatedDateTime == (time.Time{}) ||
//		b.CreatedDateTime == (time.Time{}) || b.UpdatedDateTime == (time.Time{}) {
//		log.Println("invalid merge")
//		return Player{}
//	}
//
//	c = Player{}
//	c.FirstName = mergeFirstName(&a, &b)
//	c.LastName = mergeLastName(&a, &b)
//	c.Position = mergePosition(&a, &b)
//
//	return c
//}

func mergeFirstName(a, b *Player) (firstName string) {
	firstName = ""

	if a.FirstName == b.FirstName {
		firstName = a.FirstName
	} else {
		if a.FirstName == "" {
			firstName = b.FirstName
		} else if b.FirstName == "" {
			firstName = a.FirstName
		} else {
			if a.UpdatedDateTime.After(b.UpdatedDateTime) {
				firstName = a.FirstName
			} else {
				firstName = b.FirstName
			}
		}
	}

	return firstName
}

func mergeLastName(a, b *Player) (lastName string) {
	lastName = ""

	if a.LastName == b.LastName {
		lastName = a.LastName
	} else {
		if a.LastName == "" {
			lastName = b.LastName
		} else if b.LastName == "" {
			lastName = a.LastName
		} else {
			if a.UpdatedDateTime.After(b.UpdatedDateTime) {
				lastName = a.LastName
			} else {
				lastName = b.LastName
			}
		}
	}

	return lastName
}

func mergePosition(a, b *Player) (position string) {
	position = ""

	if a.Position == b.Position {
		position = a.Position
	} else {
		if a.Position == "" {
			position = b.Position
		} else if b.Position == "" {
			position = a.Position
		} else {
			if a.UpdatedDateTime.After(b.UpdatedDateTime) {
				position = a.Position
			} else {
				position = b.Position
			}
		}
	}

	return position
}
