package model

type NotifyRequest []string

type IndexData struct {
	Code           string `ch:"code"`
    Title          string `ch:"title"`
	AdditionalInfo string `ch:"additional_info"`
}
