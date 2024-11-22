package model

import "time"

type Event struct {
	Code           string    `json:"code"            ch:"code"`
	Gender         string    `json:"gender"          ch:"gender"`
	Sport          string    `json:"sport"           ch:"sport"`
	AdditionalInfo string    `json:"additional_info" ch:"additional_info"`
	Country        string    `json:"country"         ch:"country"`
	Region         string    `json:"region"          ch:"region"`
	Locality       string    `json:"locality"        ch:"locality"`
	Stage          string    `json:"stage"           ch:"stage"`
	StartDate      time.Time `json:"start_date"      ch:"start_date"`
	EndDate        time.Time `json:"end_date"        ch:"end_date"`
	Participants   uint32    `json:"n_participants"  ch:"n_participants"`
}
