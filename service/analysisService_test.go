package service

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
	"github.com/stretchr/testify/assert"
)

func TestAnalysisLogic(t *testing.T) {
	ctx := context.Background()
	svc := New(&config.Config{})

	payload, err := os.ReadFile("../testdata/testdata.json")
	if err != nil {
		t.Fatal(err)
	}

	var analysisRequest []model.OptionsContract
	err = json.Unmarshal(payload, &analysisRequest)
	if err != nil {
		t.Fatal(err)
	}

	type test struct {
		name            string
		analysisRequest []model.OptionsContract
		checkExpected   func(*testing.T, model.AnalysisResult) bool
	}

	tests := []test{
		{
			name:            "Success",
			analysisRequest: analysisRequest,
			checkExpected: func(t *testing.T, res model.AnalysisResult) bool {
				return assert.NotNil(t, res)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := svc.AnalysisLogic(ctx, tc.analysisRequest)
			tc.checkExpected(t, res)
		})
	}
}
