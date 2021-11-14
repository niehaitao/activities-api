package model

type Activity struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

type Summary struct {
	Create int `json:"create"`
	Update int `json:"update"`
	Delete int `json:"delete"`
}
