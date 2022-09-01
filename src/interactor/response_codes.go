package interactor

import (
	"net/http"
)

const (
	OK                                  = "N200"
	TOKEN_AUTHENTICATION_ERROR          = "E300"
	NO_PERMISSION                       = "E302"
	INTERNAL_SERVER_ERROR               = "E303"
	REQUEST_DATA_EMPTY                  = "E311"
	INVALID_REQUEST_QUERIES             = "E312"
	INVALID_REQUEST_DATAS               = "E313"
	ERROR_COMMUNICATE_INTERNAL_DATABASE = "E320"
)

var ResultMessages = map[string]string{
	OK:                                  "OK.",
	TOKEN_AUTHENTICATION_ERROR:          "Token authentication error.",
	NO_PERMISSION:                       "No permission",
	INTERNAL_SERVER_ERROR:               "Internal Server Error Occured.",
	REQUEST_DATA_EMPTY:                  "Some of request data is empty.",
	INVALID_REQUEST_QUERIES:             "Please check format of request queries.",
	INVALID_REQUEST_DATAS:               "Please check request datas.",
	ERROR_COMMUNICATE_INTERNAL_DATABASE: "Can't communicate with internal database.",
}

var HTTPStatusCodes = map[string]int{
	OK:                                  http.StatusOK,
	TOKEN_AUTHENTICATION_ERROR:          http.StatusUnauthorized,
	NO_PERMISSION:                       http.StatusUnauthorized,
	INTERNAL_SERVER_ERROR:               http.StatusInternalServerError,
	REQUEST_DATA_EMPTY:                  http.StatusBadRequest,
	INVALID_REQUEST_QUERIES:             http.StatusBadRequest,
	INVALID_REQUEST_DATAS:               http.StatusBadRequest,
	ERROR_COMMUNICATE_INTERNAL_DATABASE: http.StatusInternalServerError,
}

type BasicResponse struct {
	APIResponse APIResponse `json:"api_response"`
}

type APIResponse struct {
	ResultCode    string `json:"result_code"`
	ResultMessage string `json:"result_message"`
}

func GetHTTPStatusCode(code string) int {
	return HTTPStatusCodes[code]
}

func GetAPIResponse(code string) APIResponse {
	return APIResponse{
		ResultCode:    code,
		ResultMessage: ResultMessages[code],
	}
}
