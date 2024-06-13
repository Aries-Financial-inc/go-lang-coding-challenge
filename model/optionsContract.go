package model

import "time"

type OptionType string

const (
	Call OptionType = "Call"
	Put  OptionType = "Put"
)

type LongShort string

const (
	Long  LongShort = "long"
	Short LongShort = "short"
)

// Your model here
type OptionsContract struct {
	StrikePrice    float64    `json:"strike_price"`
	Type           OptionType `json:"type"`
	Bid            float64    `json:"bid"`
	Ask            float64    `json:"ask"`
	LongShort      LongShort  `json:"long_short"`
	ExpirationDate time.Time  `json:"expiration_date"`
}

type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
