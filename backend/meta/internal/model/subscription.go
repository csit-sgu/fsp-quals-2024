package model

import "github.com/google/uuid"

type Subscription struct {
	Email          string    `json:"email"           ch:"email"`
	Code           string    `json:"code"            ch:"code"`
	Gender         string    `json:"gender"          ch:"gender"`
	Confirmation   uuid.UUID `json:"-"               ch:"confirmation"`
	Age            uint32    `json:"age"             ch:"age"`
	Sport          string    `json:"sport"           ch:"sport"`
	AdditionalInfo string    `json:"additional_info" ch:"additional_info"`
	Country        string    `json:"country"         ch:"country"`
	Region         string    `json:"region"          ch:"region"`
	Locality       string    `json:"locality"        ch:"locality"`
	Stage          string    `json:"stage"           ch:"stage"`
	EventType      string    `json:"event_type"      ch:"event_type"`
	EventScale     string    `json:"event_scale"     ch:"event_scale"`
	DateRange      DateRange `json:"date_range"      ch:"start_date"`
}

type SubscriptionConfirmation struct {
	Confirmation uuid.UUID `json:"confirmation"`
}
