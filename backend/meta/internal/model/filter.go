package model

import (
	"encoding/json"
	"time"
)

const DateFormat = "2006-01-02 00:00:00.000000"

type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	tt, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	*t = CustomTime(tt)
	return nil
}

type DateRange struct {
	From CustomTime `json:"from"`
	To   CustomTime `json:"to"`
}

func (customTime CustomTime) String() string {
	return time.Time(customTime).Format("2006-01-02")
}

type Pagination struct {
	PageSize uint32 `json:"page_size"`
	Page     uint32 `json:"page"`
}

type FilterCondition struct {
	Code           string    `json:"code"            filter:"common"   ch:"code"`
	Gender         string    `json:"gender"          filter:"common"   ch:"gender"`
	Sport          string    `json:"sport"           filter:"common"   ch:"sport"`
	AdditionalInfo string    `json:"additional_info" filter:"fuzzy"    ch:"additional_info"`
	Country        string    `json:"country"         filter:"common"   ch:"country"`
	Region         string    `json:"region"          filter:"common"   ch:"region"`
	City           string    `json:"city"            filter:"common"   ch:"city"`
	Stage          string    `json:"stage"           filter:"common"   ch:"stage"`
	DateRange      DateRange `json:"date_range"      filter:"interval" ch:"start_date"`
}

type FilterRequest struct {
	Condition      FilterCondition `json:"condition"`
	RequiredFields []string        `json:"required_fields"`
	Pagination     Pagination      `json:"pagination"`
}
