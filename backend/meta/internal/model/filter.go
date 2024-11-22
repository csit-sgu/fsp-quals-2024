package model

import (
	"time"
)

type DateRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type FilterCondition struct {
	Code           string    `json:"code"            filter:"common"   ch:"code"`
	Gender         string    `json:"gender"          filter:"common"   ch:"gender"`
	Sport          string    `json:"sport"           filter:"common"   ch:"sport"`
	AdditionalInfo string    `json:"additional_info" filter:"fuzzy"    ch:"additional_info"`
	Country        string    `json:"country"         filter:"common"   ch:"country"`
	Region         string    `json:"region"          filter:"join"     ch:"region"`
	City           string    `json:"city"            filter:"join"     ch:"city"`
	Stage          string    `json:"stage"           filter:"common"   ch:"stage"`
	DateRange      DateRange `json:"date_range"      filter:"interval" ch:"start_date"`
}

type FilterRequest struct {
	Condition      FilterCondition `json:"condition"`
	RequiredFields []string        `json:"required_fields"`
}
