package helper

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type HelperResponse struct {
	SuccessResponse func()
}

type HelperResponseStruct struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type HelperResposeRequestStruct struct {
	Status  bool     `json:"status"`
	Message []string `json:"messsage"`
}

type HelperResponseWithDataStruct struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HelperResponsePaginateStruct struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Query   interface{} `json:"query"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Link    interface{} `json:"link,omitempty"`
}

func SuccessResponse(model string, action string) HelperResponseStruct {

	return HelperResponseStruct{
		Status:  true,
		Message: typeAction(action) + model,
	}
}

func SuccessResponseWithData(model string, action string, data interface{}) HelperResponseWithDataStruct {

	return HelperResponseWithDataStruct{
		Status:  true,
		Message: typeAction(action) + model,
		Data:    data,
	}
}

func SuccessResponsePaginate(model string, action string, meta *Pagination, query *PaginationQuery, link *Pages, data interface{}) HelperResponsePaginateStruct {

	return HelperResponsePaginateStruct{
		Status:  true,
		Message: typeAction(action) + model,
		Meta:    meta,
		Query:   query,
		Link:    link,
		Data:    data,
	}
}

func ErrorResponse(err error) HelperResponseStruct {
	return HelperResponseStruct{
		Status:  false,
		Message: err.Error(),
	}
}

func ErrorNoTokenResponse() HelperResponseStruct {
	return HelperResponseStruct{
		Status:  false,
		Message: "No Token Provided",
	}
}

func ErrorNoValidTokenResponse() HelperResponseStruct {
	return HelperResponseStruct{
		Status:  false,
		Message: "Your Token is not Valid",
	}
}

func ErrorRequestResponse(err error) HelperResposeRequestStruct {
	errorMessages := []string{}
	log.Print(err)

	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		errorMessages = append(errorMessages, errorMessage)
	}

	return HelperResposeRequestStruct{
		Status:  false,
		Message: errorMessages,
	}
}

func NotFoundResponse(module string) HelperResponseStruct {
	module = strings.Title(module)

	return HelperResponseStruct{
		Status:  false,
		Message: fmt.Sprintf("Data %s ID Not Found", module),
	}
}

func ErrorCheckPassword() HelperResponseStruct {
	return HelperResponseStruct{
		Status:  false,
		Message: "Current password doesn't match on our record.",
	}
}

func typeAction(action string) string {
	message := ""

	switch action {
	case "get":
		message = "Succesfully get "
	case "create":
		message = "Successfully create "
	case "update":
		message = "Successfully update "
	case "delete":
		message = "Successfully delete "
	}

	return message
}
