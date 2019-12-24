package main

type Player struct {
  id int `json:"id" yaml:"id"`
  firstName string `json:"first_name" yaml:"firstName"`
  lastName string `json:"last_name" yaml:"firstName"`
  position string `json:"position" yaml:"position"`
  data Data `json:"data" yaml:"data"`
}
