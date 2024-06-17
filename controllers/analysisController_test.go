package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/mocks"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAnalysisHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := mocks.NewService(t)
	ctrl := new(controller)
	ctrl.s = mockService

	validPayload := []model.OptionsContract{
		{
			Type:           "Call",
			StrikePrice:    100,
			Bid:            10.05,
			Ask:            12.04,
			ExpirationDate: time.Now().AddDate(1, 0, 0),
			LongShort:      "long",
		},
	}

	invalidPayload := []model.OptionsContract{
		{
			Type:           "InvalidType", // Invalid type
			StrikePrice:    100,
			Bid:            10.05,
			Ask:            12.04,
			ExpirationDate: time.Now().AddDate(1, 0, 0),
			LongShort:      "long",
		},
	}

	type test struct {
		name          string
		payload       interface{}
		mockCall      *mock.Call
		checkExpected func(*testing.T, *httptest.ResponseRecorder)
	}

	tests := []test{
		{
			name:    "Success",
			payload: validPayload,
			mockCall: mockService.On("AnalysisLogic", mock.Anything, validPayload).
				Return(model.AnalysisResult{}).
				Once(),
			checkExpected: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "Validation Error",
			payload: invalidPayload,
			checkExpected: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "Invalid JSON",
			payload: `{"invalid_json"}`,
			checkExpected: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)

			span, ctx := opentracing.StartSpanFromContext(context.Background(), "test")
			defer span.Finish()

			jsonValue, _ := json.Marshal(tt.payload)
			c.Request = &http.Request{
				Method: "POST",
				Header: make(http.Header),
				Body:   io.NopCloser(bytes.NewBuffer(jsonValue)),
			}
			c.Request = c.Request.WithContext(ctx)

			ctrl.AnalysisHandler(c)

			tt.checkExpected(t, recorder)
		})
	}
}
