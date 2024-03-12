package model

type Email struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Pass     string `json:"pass"`
	Tags     string `json:"tags"`
	Assigned string `json:"assigned"`
	Read     string `json:"read"`
	Method   string `json:"method"`
}
