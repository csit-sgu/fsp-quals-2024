package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	Email          string    `json:"email"                     ch:"email"`
	Code           string    `json:"code,omitempty"            ch:"code"`
	Gender         string    `json:"gender,omitempty"          ch:"gender"`
	Confirmation   uuid.UUID `json:"-"                         ch:"confirmation"`
	Age            uint32    `json:"age,omitempty"             ch:"age"`
	Sport          string    `json:"sport,omitempty"           ch:"sport"`
	AdditionalInfo string    `json:"additional_info,omitempty" ch:"additional_info"`
	Country        string    `json:"country,omitempty"         ch:"country"`
	Region         string    `json:"region,omitempty"          ch:"region"`
	Locality       string    `json:"locality,omitempty"        ch:"locality"`
	Stage          string    `json:"stage,omitempty"           ch:"stage"`
	EventType      string    `json:"event_type,omitempty"      ch:"event_type"`
	EventScale     string    `json:"event_scale,omitempty"     ch:"event_scale"`
	StartDate      time.Time `json:"start_date,omitempty"      ch:"start_date"`
	EndDate        time.Time `json:"end_date,omitempty"        ch:"end_date"`
}

type SubscriptionConfirmation struct {
	Confirmation uuid.UUID `json:"confirmation"`
}
