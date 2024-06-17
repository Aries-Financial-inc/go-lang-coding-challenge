package main

import (
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/controllers"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/log"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/routes"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/service"
)

func main() {
	cfg := config.GetConfig()

	log.Info("Starting Risk Reward server::", cfg.GetString("server_port"))

	// Router, ctrl, svc and server intialization
	riskRewardServer := routes.NewRiskRewardServer()
	svc := service.New(cfg)
	ctrl := controllers.New(svc, cfg)

	// Register handlers
	riskRewardServer.SetupRouter(ctrl)

	r := riskRewardServer.Router
	// Listen and serve
	if err := http.ListenAndServe(":"+cfg.GetString("server_port"), r); err != nil {
		log.Fatal("http.ListenAndServe Failed")
	}
}
