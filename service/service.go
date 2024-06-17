package service

import "github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"

// This is the service interface
type (
	Service interface {
		IAnalysisService
	}

	service struct {
		config *config.Config
	}
)

// New creates a new instance of the Service interface.
//
// Parameters:
// - config: a pointer to a config.Config struct representing the configuration.
//
// Returns:
// - Service: a pointer to a service struct implementing the Service interface.
func New(config *config.Config) Service {
	return &service{
		config: config,
	}
}
