package exception

import "net/http"

const (
	OtherPlaceLoggedIn   = 50010
	OtherIPLoggedIn      = 50011
	OtherClientsLoggedIn = 50012
	SessionTerminated    = 50013
	AccessTokenExpired   = 50014
	RefreshTokenExpired  = 50015
	AccessTokenIllegal   = 50016
	RefreshTokenIllegal  = 50017
	VerifyCodeRequired   = 50018
	PasswordExpired      = 50019

	Unauthorized        = http.StatusUnauthorized
	BadRequest          = http.StatusBadRequest
	InternalServerError = http.StatusInternalServerError
	Forbidden           = http.StatusForbidden
	NotFound            = http.StatusNotFound
	Conflict            = http.StatusConflict

	UnKnownException = 99999
)

var (
	reasonMap = map[int]string{
		Unauthorized:         "Authentication failed",
		NotFound:             "Resource not found",
		Conflict:             "Resource already exists",
		BadRequest:           "Request is illegal",
		InternalServerError:  "Internal System Error",
		Forbidden:            "Unauthorized access",
		UnKnownException:     "Unknown exception",
		AccessTokenIllegal:   "Invalid access token",
		RefreshTokenIllegal:  "Invalid refresh token",
		OtherPlaceLoggedIn:   "Remote login",
		OtherIPLoggedIn:      "Abnormal IP login",
		OtherClientsLoggedIn: "The user has logged in through the other end",
		SessionTerminated:    "End of session",
		AccessTokenExpired:   "Visit expired, please refresh",
		RefreshTokenExpired:  "Refresh expired, please log in",
		VerifyCodeRequired:   "Abnormal operation, verification code is required for secondary confirmation",
		PasswordExpired:      "Password expired, please retrieve the password or contact the administrator to reset",
	}
)

func codeReason(code int) string {
	v, ok := reasonMap[code]
	if !ok {
		v = reasonMap[UnKnownException]
	}

	return v
}
