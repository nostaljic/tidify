package interactor

import (
	"net/http"
)

const (
	OK                                  = "N200"
	TOKEN_AUTHENTICATION_ERROR          = "E300"
	INTERNAL_SERVER_ERROR               = "E500"
	ERROR_CHECK_NESSESARY_INFORMATIONS  = "E540"
	REQUEST_DATA_EMPTY                  = "E541"
	CANNOT_LOAD_REQUESTED_DATAS         = "E542"
	INVALID_REQUEST_QUERIES             = "E544"
	INVALID_REQUEST_DATAS               = "E545"
	INVALID_REQUEST_SEQUENCE            = "E546"
	ERROR_COMMUNICATE_INTERNAL_DATABASE = "E580"
)

var ResultMessages = map[string]string{
	OK:                                  "OK.",
	TOKEN_AUTHENTICATION_ERROR:          "Authentication error with given token.",
	INTERNAL_SERVER_ERROR:               "Internal Server Error Occured.",
	ERROR_CHECK_NESSESARY_INFORMATIONS:  "Please check nessesary informations.",
	REQUEST_DATA_EMPTY:                  "Some of request data is empty.",
	CANNOT_LOAD_REQUESTED_DATAS:         "Can't load requested datas.",
	INVALID_REQUEST_QUERIES:             "Please check format of request queries.",
	INVALID_REQUEST_DATAS:               "Please check request datas.",
	INVALID_REQUEST_SEQUENCE:            "Please check request sequence.",
	ERROR_COMMUNICATE_INTERNAL_DATABASE: "Can't communicate with internal database.",
}

var HTTPStatusCodes = map[string]int{
	OK:                                  http.StatusOK,
	TOKEN_AUTHENTICATION_ERROR:          http.StatusUnauthorized,
	ERROR_CHECK_NESSESARY_INFORMATIONS:  http.StatusBadRequest,
	INTERNAL_SERVER_ERROR:               http.StatusInternalServerError,
	REQUEST_DATA_EMPTY:                  http.StatusBadRequest,
	CANNOT_LOAD_REQUESTED_DATAS:         http.StatusBadRequest,
	INVALID_REQUEST_QUERIES:             http.StatusBadRequest,
	INVALID_REQUEST_DATAS:               http.StatusBadRequest,
	INVALID_REQUEST_SEQUENCE:            http.StatusBadRequest,
	ERROR_COMMUNICATE_INTERNAL_DATABASE: http.StatusInternalServerError,
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
