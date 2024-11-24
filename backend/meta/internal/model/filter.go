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

func (t CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02"))
}

type DateRange struct {
	From CustomTime `json:"from"`
	To   CustomTime `json:"to"`
}

func (customTime CustomTime) String() string {
	return time.Time(customTime).Format("2006-01-02")
}

type Pagination struct {
	PageSize uint64 `json:"page_size"`
	Page     uint64 `json:"page"`
}

type FilterCondition struct {
	Code           string    `json:"code"            filter:"common"   ch:"code"`
	Gender         string    `json:"gender"          filter:"common"   ch:"gender"`
	Age            uint32    `json:"age"             filter:"inside"   ch:"age"`
	Sport          string    `json:"sport"           filter:"common"   ch:"sport"`
	Title          string    `json:"title"           filter:"fuzzy"    ch:"title"`
	AdditionalInfo string    `json:"additional_info" filter:"fuzzy"    ch:"additional_info"`
	Country        string    `json:"country"         filter:"common"   ch:"country"`
	Region         string    `json:"region"          filter:"common"   ch:"region"`
	Locality       string    `json:"locality"        filter:"common"   ch:"locality"`
	EventType      string    `json:"event_type"      filter:"common"   ch:"event_type"`
	EventScale     string    `json:"event_scale"     filter:"common"   ch:"event_scale"`
	DateRange      DateRange `json:"date_range"      filter:"interval" ch:"start_date"`
}

type FilterRequest struct {
	Condition      FilterCondition `json:"condition"`
	RequiredFields []string        `json:"required_fields"`
	Pagination     Pagination      `json:"pagination"`
}

type FilterResponse struct {
	Events []*Event `json:"events"`
	Total  uint64   `json:"total"`
}

type FilterView struct {
	Events []EventView
	Total  uint32
}
