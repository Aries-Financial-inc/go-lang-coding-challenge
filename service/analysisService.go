package service

import (
	"context"
	"math"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/constants"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/log"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/utils"
)

type IAnalysisService interface {
	AnalysisLogic(ctx context.Context, contracts []model.OptionsContract) model.AnalysisResult
}

// AnalysisLogic calculates the risk and reward graph, maximum profit, maximum loss, and break-even points for a set of options contracts.
//
// Parameters:
// - ctx: the context.Context object for handling cancellation and timeouts.
// - contracts: a slice of model.OptionsContract representing the options contracts.
//
// Returns:
// - model.AnalysisResult: a structure containing the risk and reward graph, maximum profit, maximum loss, and break-even points.
func (s *service) AnalysisLogic(ctx context.Context, contracts []model.OptionsContract) model.AnalysisResult {

	log := log.GetLogger(ctx)
	log.Info("Service::AnalysisLogic")

	graph := calculateXYValues(ctx, contracts)

	maxProfit := calculateMaxProfit(ctx, graph)
	maxLoss := calculateMaxLoss(ctx, graph)
	breakEvenPoints := calculateBreakEvenPoints(ctx, graph)

	return model.AnalysisResult{
		GraphData:       graph,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}
}

// calculateXYValues calculates the X and Y values for a risk and reward graph based on a set of options contracts.
//
// Parameters:
// - ctx: the context.Context object for handling cancellation and timeouts.
// - contracts: a slice of model.OptionsContract representing the options contracts.
//
// Returns:
// - graph: a slice of model.GraphPoint representing the X and Y values for the graph.
func calculateXYValues(ctx context.Context, contracts []model.OptionsContract) []model.GraphPoint {
	log := log.GetLogger(ctx)
	log.Info("Service::AnalysisLogic::calculateXYValues")

	var graph []model.GraphPoint
	minPrice := utils.GetEnvOrDefaultFloat64("MIN_PRICE", 0.0)
	maxPrice := utils.GetEnvOrDefaultFloat64("MAX_PRICE", 200.0)
	step := utils.GetEnvOrDefaultFloat64("STEP", 1.0)

	for underlying := minPrice; underlying <= maxPrice; underlying += step {
		totalPL := 0.0

		for _, contract := range contracts {
			pl := calculatePL(ctx, underlying, contract)
			totalPL += pl
		}

		graph = append(graph, model.GraphPoint{X: underlying, Y: utils.RoundToDecimal(totalPL, 2)})
	}

	return graph
}

// calculatePL calculates the profit or loss for an options contract based on the underlying price and the contract's type and position.
//
// Parameters:
// - underlying: the price of the underlying asset at the time of calculation.
// - contract: the options contract for which to calculate the profit or loss.
//
// Returns:
// - pl: the profit or loss for the options contract.
func calculatePL(ctx context.Context, underlying float64, contract model.OptionsContract) float64 {

	var pl float64
	switch contract.Type {
	case constants.Call:
		if contract.LongShort == constants.Long {
			pl = math.Max(0, underlying-contract.StrikePrice) - contract.Ask
		} else {
			pl = contract.Bid - math.Max(0, underlying-contract.StrikePrice)
		}
	case constants.Put:
		if contract.LongShort == constants.Long {
			pl = math.Max(0, contract.StrikePrice-underlying) - contract.Ask
		} else {
			pl = contract.Bid - math.Max(0, contract.StrikePrice-underlying)
		}
	}
	return pl
}

// calculateMaxProfit calculates the maximum profit from a risk-reward graph.
//
// Parameters:
// - ctx: the context.Context object for handling cancellation and timeouts.
// - graph: a slice of model.GraphPoint representing the X and Y values for the graph.
//
// Returns:
// - maxProfit: the maximum profit value from the graph.
func calculateMaxProfit(ctx context.Context, graph []model.GraphPoint) float64 {
	log := log.GetLogger(ctx)
	log.Info("Service::AnalysisLogic::calculateMaxProfit")

	maxProfit := graph[0].Y
	for _, point := range graph {
		if point.Y > maxProfit {
			maxProfit = point.Y
		}
	}
	return maxProfit
}

// calculateMaxLoss calculates the maximum loss from a risk-reward graph.
//
// Parameters:
// - ctx: the context.Context object for handling cancellation and timeouts.
// - graph: a slice of model.GraphPoint representing the X and Y values for the graph.
//
// Returns:
// - maxLoss: the maximum loss value from the graph.
func calculateMaxLoss(ctx context.Context, graph []model.GraphPoint) float64 {
	log := log.GetLogger(ctx)
	log.Info("Service::AnalysisLogic::calculateMaxLoss")

	maxLoss := graph[0].Y
	for _, point := range graph {
		if point.Y < maxLoss {
			maxLoss = point.Y
		}
	}
	return maxLoss
}

// calculateBreakEvenPoints calculates the break-even points for a risk and reward graph.
//
// Parameters:
// - ctx: the context.Context object for handling cancellation and timeouts.
// - graph: a slice of model.GraphPoint representing the X and Y values for the graph.
//
// Returns:
// - breakEvenPoints: a slice of float64 representing the break-even points.
func calculateBreakEvenPoints(ctx context.Context, graph []model.GraphPoint) []float64 {
	log := log.GetLogger(ctx)
	log.Info("Service::AnalysisLogic::calculateBreakEvenPoints")

	var breakEvenPoints []float64
	for i := 1; i < len(graph); i++ {
		if graph[i-1].Y*graph[i].Y <= 0 {
			breakEvenPoints = append(breakEvenPoints, graph[i].X)
		}
	}

	return breakEvenPoints
}
