package request

import (
	"encoding/json"
	"fmt"
	"github.com/ducketlab/mingo/exception"
	"io/ioutil"
	"net/http"
)

var (
	// BodyMaxContentLength 64Mb
	BodyMaxContentLength int64 = 1 << 26
)

func ReadBody(r *http.Request) ([]byte, error) {
	if r.ContentLength == 0 {
		return nil, exception.NewBadRequest("request body is empty")
	}

	if r.ContentLength > BodyMaxContentLength {
		return nil, exception.NewBadRequest(
			"the body exceeding the maximum limit, max size %dM",
			BodyMaxContentLength/1024/1024)
	}

	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		return nil, exception.NewBadRequest(
			fmt.Sprintf("read request body error, %s", err))
	}

	return body, nil

}

func GetDataFromRequest(r *http.Request, v interface{}) error {
	body, err := ReadBody(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
