package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctrl := New(nil, nil)
	assert.NotNil(t, ctrl)
}
