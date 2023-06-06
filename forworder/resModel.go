package main

type BackendAResponse struct {
	Status  string `json:"status"`
	Backend string `json:"backend"`
	Version string `json:"version"`
	Header  string `json:"header"`
	AResponse
}
type BackendBResponse struct {
	Status  string `json:"status"`
	Backend string `json:"backend"`
	Version string `json:"version"`
	Header  string `json:"header"`
	BResponse
}

type BResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
	Summary   string `json:"summary"`
	Publisher string `jsson:"publisher"`
	Rating    Rating `json:"rating"`
}

type Rating struct {
	Average     float32 `json:"average"`
	ToatalVotes int     `json:"totalVotes"`
}

type AResponse struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
	Rating float32 `json:"rating"`
}
