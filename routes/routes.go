package routes

import (
	"go-challenge/controllers"

	"github.com/gin-gonic/gin"
)

// OptionsContract structure for the request body
type OptionsContract struct {
	// Your code here
}

// AnalysisResult structure for the response body
type AnalysisResult struct {
	GraphData       []GraphPoint `json:"graph_data"`
	MaxProfit       float64      `json:"max_profit"`
	MaxLoss         float64      `json:"max_loss"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}

// GraphPoint structure for X & Y values of the risk & reward graph
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", controllers.AnalysisHandler)

	return router
}
