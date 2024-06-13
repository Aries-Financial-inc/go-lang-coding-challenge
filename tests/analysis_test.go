package tests

import (
	"bytes"
	"encoding/json"
	"go-challenge/controllers"
	"go-challenge/model"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOptionsContractModelValidation(t *testing.T) {
	testDataJson, err := os.ReadFile("../testdata/testdata.json")
	assert.NoError(t, err)

	var requestPayload []model.OptionsContract
	err = json.Unmarshal(testDataJson, &requestPayload)
	assert.NoError(t, err)
}

func TestAnalysisEndpoint(t *testing.T) {
	// Initialize test data from json
	testDataJson, err := os.ReadFile("../testdata/testdata.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var requestPayload []model.OptionsContract
	if err := json.Unmarshal(testDataJson, &requestPayload); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

	payload, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	router := gin.Default()
	router.POST("/analyze", controllers.AnalysisHandler)

	req, err := http.NewRequest("POST", "/analyze", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	// Record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response model.AnalysisResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	assert.NotNil(t, response.XYValues)
	assert.NotNil(t, response.MaxProfit)
	assert.NotNil(t, response.MaxLoss)
	assert.NotNil(t, response.BreakEvenPoints)
}

func TestIntegration(t *testing.T) {
	// Initialize test data from json
	testDataJson, err := os.ReadFile("../testdata/testdata.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var requestPayload []model.OptionsContract
	if err := json.Unmarshal(testDataJson, &requestPayload); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

	// Marshal the request payload to JSON
	payload, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	router := gin.Default()
	router.POST("/analyze", controllers.AnalysisHandler)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/analyze", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	// Record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Decode the response body
	var response model.AnalysisResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	// Assert the response contents
	assert.NotEmpty(t, response.XYValues)
	assert.NotEqual(t, 0.0, response.MaxProfit)
	assert.NotEqual(t, 0.0, response.MaxLoss)
	assert.NotEmpty(t, response.BreakEvenPoints)

	// Additional checks to ensure correctness
	assert.True(t, len(response.XYValues) > 0, "XYValues should have data points")
	assert.True(t, len(response.BreakEvenPoints) > 0, "BreakEvenPoints should have at least one point")

	// Print the results (optional, for visual confirmation during testing)
	t.Logf("XY Values: %v", response.XYValues)
	t.Logf("Max Profit: %f", response.MaxProfit)
	t.Logf("Max Loss: %f", response.MaxLoss)
	t.Logf("Break Even Points: %v", response.BreakEvenPoints)
}
