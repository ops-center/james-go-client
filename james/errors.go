package james

import (
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
	"io"
	"net/http"
)

type JamesServerError struct {
	StatusCode int    `json:"statusCode"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	Details    string `json:"details"`
}

func (e JamesServerError) Error() string {
	return fmt.Sprintf("james server responded with status code: %d, type: %s, message: %s, details: %s\n", e.StatusCode, e.Type, e.Message, e.Details)
}

func toJamesServerError(r *http.Response, errorList ...error) *JamesServerError {
	if r == nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    errors.Join(errorList...).Error(),
		}
	}

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    errors.Join(append(errorList, err)...).Error(),
		}
	}

	return &JamesServerError{
		StatusCode: r.StatusCode,
		Message:    errors.Join(append(errorList, errors.New(string(bodyData[:])))...).Error(),
	}
}

func (j JamesServerError) GetError() error {
	err := fmt.Errorf("%s", j.Type)
	if !isEmptyString(j.Message) {
		err = errors2.Wrap(err, j.Message)
	}
	if !isEmptyString(j.Details) {
		err = errors2.Wrap(err, j.Details)
	}

	return err
}

func (j JamesServerError) GetStatusCode() int {
	return j.StatusCode
}

func (j JamesServerError) GetMessage() string {
	return j.Message
}
