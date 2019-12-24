package main

type Meta struct {
  totalPages int `json:"total_pages" yaml:"totalPages"`
  currentPage int `json:"current_page" yaml:"currentPage`
  nextPage int `json:"next_page" yaml"nextPage"`
  perPage int  `json:"per_page" yaml:"perPage"`
  totalCount int  `json:"total_count" yaml:"totalCount`
}

