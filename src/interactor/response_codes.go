package interactor

import (
	"net/http"
)

const (
	OK                                 = "N100"
	TOKEN_AUTHENTICATION_ERROR         = "E300"
	ERROR_CHECK_NESSESARY_INFORMATIONS = "E530"
	REQUEST_DATA_EMPTY                 = "E531"
	CANNOT_LOAD_REQUESTED_DATAS        = "E532"
	ERROR_SAVE_NESSESARY_INFORMATIONS  = "E533"
	INVALID_REQUEST_QUERIES            = "E534"
	INVALID_REQUEST_DATAS              = "E535"
	INVALID_REQUEST_SEQUENCE           = "E536"
	INVALID_MODULE_ACCESS              = "E537"
)

var ResultMessages = map[string]string{
	OK:                                 "OK.",
	TOKEN_AUTHENTICATION_ERROR:         "Authentication error with given token.",
	ERROR_CHECK_NESSESARY_INFORMATIONS: "Error when check usage of nessesary informations.",
	REQUEST_DATA_EMPTY:                 "Some request data empty.",
	CANNOT_LOAD_REQUESTED_DATAS:        "Cannot load requested datas.",
	ERROR_SAVE_NESSESARY_INFORMATIONS:  "Error when saving nessesary informations.",
	INVALID_REQUEST_QUERIES:            "Invalid format of request queries.",
	INVALID_REQUEST_DATAS:              "Invalid request datas.",
	INVALID_REQUEST_SEQUENCE:           "Invalid request sequence.",
	INVALID_MODULE_ACCESS:              "Invalid module access",
}

var HTTPStatusCodes = map[string]int{
	OK:                                 http.StatusOK,
	TOKEN_AUTHENTICATION_ERROR:         http.StatusBadRequest,
	ERROR_CHECK_NESSESARY_INFORMATIONS: http.StatusBadRequest,
	REQUEST_DATA_EMPTY:                 http.StatusBadRequest,
	CANNOT_LOAD_REQUESTED_DATAS:        http.StatusBadRequest,
	ERROR_SAVE_NESSESARY_INFORMATIONS:  http.StatusInternalServerError,
	INVALID_REQUEST_QUERIES:            http.StatusBadRequest,
	INVALID_REQUEST_DATAS:              http.StatusBadRequest,
	INVALID_REQUEST_SEQUENCE:           http.StatusBadRequest,
	INVALID_MODULE_ACCESS:              http.StatusBadRequest,
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
