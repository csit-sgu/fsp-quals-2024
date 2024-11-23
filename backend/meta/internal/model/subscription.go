package model

type Subscription struct {
	Email  string          `json:"email"  ch:"email"`
	Filter FilterCondition `json:"filter" ch:"email"`
}
