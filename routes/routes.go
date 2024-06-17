package routes

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/controllers"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/middleware"
	"github.com/gin-gonic/gin"
)

type RiskRewardServer struct {
	Router *gin.Engine
}

// NewRiskRewardServer creates a new instance of RiskRewardServer and initializes its Router field with a default Gin engine.
//
// Returns:
// - *RiskRewardServer: A pointer to the newly created RiskRewardServer instance.
func NewRiskRewardServer() *RiskRewardServer {
	server := new(RiskRewardServer)
	server.Router = gin.Default()
	return server
}

// SetupRouter sets up the routes for the RiskRewardServer.
//
// Parameters:
// - c: an instance of controllers.Controller.
// - middlewares: variadic list of gin.HandlerFunc for middleware.
func (s *RiskRewardServer) SetupRouter(c controllers.Controller, middlewares ...gin.HandlerFunc) {
	group := s.Router.Group("/",
		middleware.TransactionInMiddleware(), // transaction middleware
		middleware.RequestLogger(),           // logger middleware
	) // group for all routes and middlewares. Here the global middlewares are applied for all route groups
	group.Use(middlewares...)

	group.POST("/analyze", c.AnalysisHandler)

}
