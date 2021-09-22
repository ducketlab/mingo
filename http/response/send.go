package response

import (
	"encoding/json"
	"fmt"
	"github.com/ducketlab/mingo/exception"
	"net/http"
)

func Failed(w http.ResponseWriter, err error) {
	var (
		errCode  int
		httpCode int
		ns       string
		reason   string
		data     interface{}
		meta     interface{}
	)

	switch t := err.(type) {
	case exception.APIException:
		errCode = t.ErrorCode()
		reason = t.Reason()
		data = t.Data()
		meta = t.Meta()
		ns = t.Namespace()
	default:
		errCode = exception.UnKnownException
	}

	if errCode/100 >= 1 && errCode/100 <= 5 {
		httpCode = errCode
	} else {
		httpCode = http.StatusOK
	}

	resp := Data{
		Code:      &errCode,
		Namespace: ns,
		Reason:    reason,
		Message:   err.Error(),
		Data:      data,
		Meta:      meta,
	}

	// set response headers
	w.Header().Set("Content-Type", "application/json")

	// if marshal json error, use string to response
	respByt, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMSG := fmt.Sprintf(`{"status":"error", "message": "encoding to json error, %s"}`, err)
		w.Write([]byte(errMSG))
		return
	}

	w.WriteHeader(httpCode)
	w.Write(respByt)
	return
}

// Success use to response success data
func Success(w http.ResponseWriter, data interface{}) {
	c := 0
	resp := Data{
		Code:    &c,
		Message: "",
		Data:    data,
	}

	// set response headers
	w.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMSG := fmt.Sprintf(`{"status":"error", "message": "encoding to json error, %s"}`, err)
		w.Write([]byte(errMSG))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}
