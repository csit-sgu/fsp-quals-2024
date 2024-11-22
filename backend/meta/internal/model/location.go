package model

type Country struct {
	Country string `json:"country" ch:"country"`
}

type Region struct {
	Region string `json:"region" ch:"region"`
}

type Locality struct {
	Locality string `json:"locality" ch:"locality"`
}

type Sport struct {
	Sport string `json:"sport" ch:"sport"`
}
