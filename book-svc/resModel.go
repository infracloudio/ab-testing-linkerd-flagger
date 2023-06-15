package main

type BackendAResponse struct {
	Status  string `json:"status"`
	Backend string `json:"backend"`
	Version string `json:"version"`
	Header  string `json:"header"`
	AResponse
}

type AResponse struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
	Rating float32 `json:"rating"`
}
