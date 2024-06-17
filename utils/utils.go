package utils

import (
	"math"
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/config"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/constants"
)

var (
	RR1001 = setErrorFields(http.StatusBadRequest, constants.RR1001, "BAD REQUEST")
)

// setErrorFields generates a map with HTTP_STATUS, CODE, and ERR_MESSAGE fields.
//
// Parameters:
// - httpStatus: an integer representing the HTTP status.
// - code: a string representing the error code.
// - errMessage: a string representing the error message.
// Returns:
// - map[string]interface{}: a map with HTTP_STATUS, CODE, and ERR_MESSAGE fields.
func setErrorFields(httpStatus int, code string, errMessage string) map[string]interface{} {
	return map[string]interface{}{
		"HTTP_STATUS": httpStatus,
		"CODE":        code,
		"ERR_MESSAGE": errMessage,
	}
}

// CustomErrorFields updates the "ERR_MESSAGE" field in the given map with the provided custom message.
//
// Parameters:
// - setErrorFields: a map[string]interface{} representing the error fields.
// - customMessage: a string representing the custom error message.
//
// Returns:
// - map[string]interface{}: the updated error fields map with the custom error message.
func CustomErrorFields(setErrorFields map[string]interface{}, customMessage string) map[string]interface{} {
	setErrorFields["ERR_MESSAGE"] = customMessage
	return setErrorFields
}

// GetEnvOrDefaultFloat64 retrieves the value of the specified key from the configuration and returns it as a float64.
// If the value is 0, it returns the defaultValue instead.
//
// Parameters:
// - key: the key to retrieve the value for.
// - defaultValue: the default value to return if the retrieved value is 0.
//
// Returns:
// - float64: the retrieved value or the defaultValue if the retrieved value is 0.
func GetEnvOrDefaultFloat64(key string, defaultValue float64) float64 {
	cfg := config.GetConfig()
	value := cfg.GetFloat64(key)
	if value == 0 {
		value = defaultValue
	}
	return value
}

// RoundToDecimal rounds a float64 number to a specified number of decimal places.
//
// Parameters:
// - num: the float64 number to be rounded.
// - decimalPlaces: the number of decimal places to round to.
//
// Returns:
// - float64: the rounded number.
func RoundToDecimal(num float64, decimalPlaces int) float64 {
	precision := math.Pow10(decimalPlaces)
	return math.Round(num*precision) / precision
}
