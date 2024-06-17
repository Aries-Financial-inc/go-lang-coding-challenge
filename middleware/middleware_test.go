package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type test struct {
		name    string
		request *http.Request
		code    int
	}

	tests := []test{
		{
			name:    "middleware applied",
			request: httptest.NewRequest(http.MethodGet, "/", nil),
			code:    http.StatusNotFound,
		},

		{
			name:    "middleware not applied",
			request: httptest.NewRequest(http.MethodPost, "/", nil),
			code:    http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			_, engine := gin.CreateTestContext(w)
			engine.Use(RequestLogger())
			engine.ServeHTTP(w, tc.request)
			assert.Equal(t, tc.code, w.Code)
		})
	}
}

func TestGenerateCorrelationId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type test struct {
		name    string
		request *http.Request
		code    int
	}

	tests := []test{
		{
			name:    "middleware applied",
			request: httptest.NewRequest(http.MethodGet, "/", nil),
			code:    http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			_, engine := gin.CreateTestContext(w)
			engine.Use(TransactionInMiddleware())
			engine.ServeHTTP(w, tc.request)
			assert.Equal(t, tc.code, w.Code)
		})
	}
}
