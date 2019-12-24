package main

type Data struct {
  gamesPlayed int `json:"games_played" yaml:"gamesPlayed"`
  min string `json:"min" yaml:"min"`
  fg3m float64 `json:"fg3m" yaml:"fg3m"`
  reb float64 `json:"red" yaml:"red"`
  ast float64 `json:"ast" yaml:"ast"`
  stl float64 `json:"stl" yaml:"stl"`
  blk float64 `json:"blk" yaml:"blk"`
  turnover float64 `json:"turnover" yaml:"turnover"`
  pts float64 `json:"pts" yaml:"pts"`
  fgPct float64 `json:"fg_pct" yaml:"fgPct"`
  ftPct float64 `json:"ft_pct" yaml:"ftPct"`
}
