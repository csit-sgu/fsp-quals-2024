package model

import (
	"time"
)

type FilterCondition struct {
	Code           string    `json:"code"            `
	Gender         string    `json:"gender"          `
	Sport          string    `json:"sport"           `
	AdditionalInfo string    `json:"additional_info" `
	Country        string    `json:"country"         `
	Region         string    `json:"region"          `
	City           string    `json:"city"            `
	Stage          string    `json:"stage"           `
	StartDate      time.Time `json:"start_date"      `
	EndDate        time.Time `json:"end_date"        `
}

type FilterRequest struct {
	Condition      FilterCondition `json:"condition"`
	RequiredFields []string        `json:"required_fields"`
}
