package james

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
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

func unmarshalToJamesServerError(r *http.Response) *JamesServerError {
	if r == nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    "nil response",
		}
	}

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to read response body: %s", err),
		}
	}

	return &JamesServerError{
		StatusCode: r.StatusCode,
		Message:    string(bodyData[:]),
	}
}

func (j JamesServerError) GetError() error {
	err := fmt.Errorf("%s", j.Type)
	if !isEmptyString(j.Message) {
		err = errors.Wrap(err, j.Message)
	}
	if !isEmptyString(j.Details) {
		err = errors.Wrap(err, j.Details)
	}

	return err
}

func (j JamesServerError) GetStatusCode() int {
	return j.StatusCode
}

func (j JamesServerError) GetMessage() string {
	return j.Message
}
