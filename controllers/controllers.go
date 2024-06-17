package controllers

import (
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/service"
)

type (
	Controller interface {
		IAnalysisController
	}

	controller struct {
		s   service.Service
		cfg *config.Config
	}
)

func New(svc service.Service, cfg *config.Config) Controller {
	return &controller{
		s:   svc,
		cfg: cfg,
	}
}
