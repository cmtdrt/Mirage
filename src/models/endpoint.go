package models

// Endpoint represents a route that will be created
type Endpoint struct {
	Method      string  `json:"method"`
	Description *string `json:"description,omitempty"`
	Path        string  `json:"path"`
	Status      *int    `json:"status,omitempty"`
	Delay       *int    `json:"delay,omitempty"`
	Response    any     `json:"response"`
}
