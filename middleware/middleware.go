package middleware

import (
	"net/http"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/constants"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LogRecord struct {
	http.ResponseWriter
	status int
}

// TransactionInMiddleware is a middleware function that generates a correlation ID for each incoming request.
//
// It takes a gin.Context object as a parameter and returns a gin.HandlerFunc.
// The gin.HandlerFunc is responsible for generating a correlation ID using the GenerateCorrelationId function
// and then calling the next handler in the chain using the c.Next() method.
func TransactionInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		GenerateCorrelationId()
		c.Next()
	}
}

// GenerateCorrelationId generates a new UUID and assigns it to the CorrelationId field in the constants package.
func GenerateCorrelationId() {
	constants.CorrelationId = uuid.New().String()
}

// Request Logger Middleware to log requests and responses
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := log.GetLogger(c)

		log.WithContext(c).Infof("Request: method: %s, url: %s", c.Request.Method, c.Request.RequestURI)
		record := &LogRecord{
			ResponseWriter: &LogRecord{c.Writer, http.StatusOK},
			status:         http.StatusOK,
		}

		log.WithContext(c).Infof("Incoming request: method: %s, url: %s", c.Request.Method, c.Request.RequestURI)

		c.Next()

		record.status = c.Writer.Status()
		log.WithContext(c).Infof("Response: method: %s, url: %s, status: %d", c.Request.Method, c.Request.RequestURI, record.status)

	}

}
