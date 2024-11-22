package model

import (
	"time"
)

type FilterCondition struct {
	Code           string    `json:"code"            db:"code"`
	Gender         string    `json:"gender"          db:"gender"`
	Sport          string    `json:"sport"           db:"sport"`
	AdditionalInfo string    `json:"additional_info" db:"additional_info"`
	Country        string    `json:"country"         db:"country"`
	Region         string    `json:"region"          db:"region"`
	City           string    `json:"city"            db:"city"`
	Stage          string    `json:"stage"           db:"stage"`
	StartDate      time.Time `json:"start_date"      db:"start_date"`
	EndDate        time.Time `json:"end_date"        db:"end_date"`
}
