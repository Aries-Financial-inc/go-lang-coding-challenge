package model

import "time"

type OptionsContract struct {
	Type           string    `json:"type" validate:"required,oneof=Call Put"`
	StrikePrice    float64   `json:"strike_price" validate:"required"`
	Bid            float64   `json:"bid" validate:"required"`
	Ask            float64   `json:"ask" validate:"required"`
	ExpirationDate time.Time `json:"expiration_date" validate:"required"`
	LongShort      string    `json:"long_short" validate:"required,oneof=long short"`
}
