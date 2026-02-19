package models

// Input represents the content of the JSON configuration file
type Input struct {
	Endpoints []Endpoint `json:"endpoints"`
}
