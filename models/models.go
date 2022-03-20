package models

type Path struct {
	Path string `json:"path"`
}

type PathsList struct {
	Paths []*Path `json:"paths"`
	Count int64   `json:"count"`
}
