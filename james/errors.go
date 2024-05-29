package james

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type JamesServerError struct {
	StatusCode int    `json:"statusCode"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	Details    string `json:"details"`
}

func newServerError(resp *http.Response, err error) *JamesServerError {
	if resp == nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    errors.Wrap(err, "failed to read response body").Error(),
		}
	}

	return &JamesServerError{
		StatusCode: resp.StatusCode,
		Message:    errors.Wrap(err, string(body)).Error(),
	}
}

func (j *JamesServerError) Error() string {
	return j.Message
}

func (j *JamesServerError) GetError() error {
	err := fmt.Errorf("%s", j.Type)
	if !isEmptyString(j.Message) {
		err = errors.Wrap(err, j.Message)
	}
	if !isEmptyString(j.Details) {
		err = errors.Wrap(err, j.Details)
	}

	return err
}
