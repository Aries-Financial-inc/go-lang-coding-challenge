package service

import (
	"testing"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	svc := New(&config.Config{})
	assert.NotNil(t, svc)
}
