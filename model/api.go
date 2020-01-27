package model

type Api struct {
	Players []Player `json:"players"`
	Statistics []Player `json:"statistics"`
}
