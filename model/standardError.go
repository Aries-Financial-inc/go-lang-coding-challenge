package model

type StandardError struct {
	Version        string             `json:"version"`
	HttpStatusCode int                `json:"httpStatusCode"`
	Errors         []APIErrorResponse `json:"errors"`
}

type APIErrorResponse struct {
	Code              string `json:"code"`
	Message           string `json:"message"`
	AdditionalMessage string `json:"additionalMessage"`
}
