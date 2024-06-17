package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	ctrlMock "github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/mocks"
	"github.com/gin-gonic/gin"
)

func TestSetupRouter(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic handled")
		}
	}()
	gin.SetMode(gin.TestMode)
	ctrl := ctrlMock.NewController(t)

	server := NewRiskRewardServer()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	server.SetupRouter(ctrl)
}
