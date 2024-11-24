package model

import "time"

type LocationData struct {
	Country  string `json:"country"  ch:"country"`
	Region   string `json:"region"   ch:"region"`
	Locality string `json:"locality" ch:"locality"`
}

type AgeData struct {
	Gender     string `json:"gender"      ch:"gender"`
	LeftBound  uint32 `json:"left_bound"  ch:"left_bound"`
	RightBound uint32 `json:"right_bound" ch:"right_bound"`
	Original   string `json:"original"    ch:"extra_mapping"`
}

type EventView struct {
	Code           string    `json:"code"            ch:"code"`
	StartDate      time.Time `json:"start_date"      ch:"start_date"`
	Country        string    `json:"country"         ch:"country"`
	Region         string    `json:"region"          ch:"region"`
	Locality       string    `json:"locality"        ch:"locality"`
	Gender         string    `json:"gender"          ch:"gender"`
	LeftBound      uint32    `json:"left_bound"      ch:"left_bound"`
	RightBound     uint32    `json:"right_bound"     ch:"right_bound"`
	Title          string    `json:"title"           ch:"title"`
	AdditionalInfo string    `json:"additional_info" ch:"additional_info"`
	Participants   uint32    `json:"n_participants"  ch:"n_participants"`
	EndDate        time.Time `json:"end_date"        ch:"end_date"`
	Sport          string    `json:"sport"           ch:"sport"`
	ExtraMapping   string    `json:"extra_mapping"   ch:"extra_mapping"`
	EventScale     string    `json:"event_scale"     ch:"event_scale"`
	EventType      string    `json:"event_type"      ch:"event_type"`
	PageIndex      uint64    `json:"page_index"      ch:"page_index"`
}

type Event struct {
	Code           string         `json:"code"            ch:"code"`
	StartDate      CustomTime     `json:"start_date"      ch:"start_date"`
	LocationData   []LocationData `json:"location_data"   ch:"location_data"`
	AgeData        []AgeData      `json:"age_data"        ch:"age_data"`
	Title          string         `json:"title"           ch:"title"`
	AdditionalInfo string         `json:"additional_info" ch:"additional_info"`
	Participants   uint32         `json:"n_participants"  ch:"n_participants"`
	Stage          string         `json:"stage"           ch:"stage"`
	EndDate        CustomTime     `json:"end_date"        ch:"end_date"`
	Sport          string         `json:"sport"           ch:"sport"`
}
