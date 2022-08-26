package auth

import (
	"net/http"
)

const (
	OK                         = "N200"
	TOKEN_AUTHENTICATION_ERROR = "E300"
	TOKEN_EXPIRED              = "E301"
)

var ResultMessages = map[string]string{
	OK:                         "OK.",
	TOKEN_AUTHENTICATION_ERROR: "Authentication error with given token.",
	TOKEN_EXPIRED:              "Already expired token",
}

var HTTPStatusCodes = map[string]int{
	OK:                         http.StatusOK,
	TOKEN_AUTHENTICATION_ERROR: http.StatusUnauthorized,
	TOKEN_EXPIRED:              http.StatusUnauthorized,
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
