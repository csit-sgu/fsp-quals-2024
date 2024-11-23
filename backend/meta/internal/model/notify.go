package model

type NotifyRequest []string

type IndexData struct {
	Code           string `json:"code"`
	Title          string `json:"title"`
	AdditionalInfo string `json:"additional_info"`
}
