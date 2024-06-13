package controllers

import (
	"go-challenge/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnalysisHandler(c *gin.Context) {
	var contracts []model.OptionsContract

	if err := c.ShouldBindJSON(&contracts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(contracts) > 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Overflow of contracts"})
		return
	}

	response := performAnalysis(contracts)

	c.JSON(http.StatusOK, response)
}

func performAnalysis(contracts []model.OptionsContract) model.AnalysisResponse {
	underlyingPrices := generateUnderlyingPrices()
	var xyValues []model.XYValue
	maxProfit, maxLoss := calculateMaxProfitLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	for _, price := range underlyingPrices {
		profitLoss := calculateProfitLoss(contracts, price)
		xyValues = append(xyValues, model.XYValue{price, profitLoss})
	}

	return model.AnalysisResponse{
		XYValues:        xyValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}
}

func generateUnderlyingPrices() []float64 {
	// Generate a range of underlying prices for X values
	var prices []float64
	for i := 0; i <= 200; i++ {
		prices = append(prices, float64(i))
	}
	return prices
}

func calculateMaxProfitLoss(contracts []model.OptionsContract) (float64, float64) {
	maxProfit := -1.0
	maxLoss := 0.0

	prices := generateUnderlyingPrices()
	for _, price := range prices {
		profitLoss := calculateProfitLoss(contracts, price)
		if maxProfit == -1.0 || profitLoss > maxProfit {
			maxProfit = profitLoss
		}
		if profitLoss < maxLoss {
			maxLoss = profitLoss
		}
	}

	return maxProfit, maxLoss
}

func calculateBreakEvenPoints(contracts []model.OptionsContract) []float64 {
	var breakEvenPoints []float64
	prices := generateUnderlyingPrices()
	previousProfitLoss := calculateProfitLoss(contracts, prices[0])

	for _, price := range prices[1:] {
		currentProfitLoss := calculateProfitLoss(contracts, price)
		if previousProfitLoss*currentProfitLoss <= 0 {
			breakEvenPoints = append(breakEvenPoints, price)
		}
		previousProfitLoss = currentProfitLoss
	}

	return breakEvenPoints
}

func calculateProfitLoss(contracts []model.OptionsContract, underlyingPrice float64) float64 {
	totalProfitLoss := 0.0

	for _, contract := range contracts {
		var profitLoss float64

		switch contract.Type {
		case model.Call:
			if contract.LongShort == model.Long {
				profitLoss = max(0, underlyingPrice-contract.StrikePrice) - contract.Ask
			} else {
				profitLoss = contract.Bid - max(0, underlyingPrice-contract.StrikePrice)
			}
		case model.Put:
			if contract.LongShort == model.Long {
				profitLoss = max(0, contract.StrikePrice-underlyingPrice) - contract.Ask
			} else {
				profitLoss = contract.Bid - max(0, contract.StrikePrice-underlyingPrice)
			}
		}

		totalProfitLoss += profitLoss
	}

	return totalProfitLoss
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
