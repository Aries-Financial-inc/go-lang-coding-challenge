package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/constants"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/log"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/model"
	"github.com/Aries-Financial-inc/golang-dev-logic-challenge-gaurav-malwe/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type IAnalysisController interface {
	AnalysisHandler(ginCtx *gin.Context)
}

// AnalysisHandler handles the analysis for options contracts.
//
// ginCtx: *gin.Context for handling HTTP request and response.
// No return value.
func (c *controller) AnalysisHandler(ginCtx *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(ginCtx.Request.Context(), "IAnalysisController::AnalysisHandler")
	defer span.Finish()

	log := log.GetLogger(ctx)
	log.Info("Controller::AnalysisHandler")

	payload, err := validateModelRequest(ginCtx)
	if checkError(ginCtx, err) {
		return
	}

	riskRewardResult := c.s.AnalysisLogic(ctx, payload)

	ginCtx.JSON(http.StatusOK, riskRewardResult)

}

// validateModelRequest validates the request payload and returns the validated options contracts and any error encountered.
//
// Parameters:
// - ginCtx: a pointer to a gin.Context object representing the request context.
//
// Returns:
// - []model.OptionsContract: an array of validated options contracts.
// - error: an error object if validation fails, otherwise nil.
func validateModelRequest(ginCtx *gin.Context) ([]model.OptionsContract, error) {
	var payload []model.OptionsContract
	var err error
	// check binding
	if err := ginCtx.ShouldBindJSON(&payload); err != nil {
		return payload, err
	}

	validate := validator.New()

	for _, contract := range payload {
		// check validation
		err = validate.Struct(contract)
		if err != nil {
			arr := listErrors(err)
			return payload, fmt.Errorf("%#v", utils.CustomErrorFields(utils.RR1001, ("Invalid/missing input parameters: "+arr)))
		}
	}
	return payload, nil
}

func listErrors(err error) string {
	var arr []string
	for _, err := range err.(validator.ValidationErrors) {
		arr = append(arr, err.Field())
	}
	str := strings.Join(arr, ", ")
	return str
}

func checkError(ginCtx *gin.Context, err error) bool {
	log := log.GetLogger(ginCtx.Request.Context())
	if err != nil {
		log.WithContext(ginCtx.Request.Context()).Error(err)
		writeErrorOnResponse(ginCtx.Writer, utils.CustomErrorFields(utils.RR1001, err.Error()))
		return true
	}
	return false
}

func writeErrorOnResponse(responseWriter http.ResponseWriter, fields map[string]interface{}) {
	httpStatus, _ := fields["HTTP_STATUS"].(int)
	additionalMessage, _ := fields["ADDITIONAL_MESSAGE"].(string)

	response := model.StandardError{
		Version:        constants.VERSION,
		HttpStatusCode: httpStatus,
		Errors: []model.APIErrorResponse{
			{
				Code:              fields["CODE"].(string),
				Message:           fields["ERR_MESSAGE"].(string),
				AdditionalMessage: additionalMessage,
			},
		},
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(httpStatus)
	json.NewEncoder(responseWriter).Encode(response)
}
