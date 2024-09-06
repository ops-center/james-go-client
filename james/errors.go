package james

import (
	"fmt"
	"io"
	"net/http"

	"git.sr.ht/~rockorager/go-jmap"
	"github.com/pkg/errors"
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
		Message:    string(body),
	}
}

func IsUnauthorizedError(err error) bool {
	var je *JamesServerError
	if errors.As(err, &je) && je.StatusCode == 401 {
		return true
	}

	var re *jmap.RequestError
	return errors.As(err, &re) && re.Status == 401
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

func serverError(statusCode int, errType, message string) *JamesServerError {
	return &JamesServerError{
		StatusCode: statusCode,
		Type:       errType,
		Message:    message,
	}
}

var (
	ErrUserNotExists  = serverError(404, "USER_DOES_NOT_EXISTS", "User does not exist")
	ErrObjectCreation = serverError(500, "OBJECT_CREATION_ERROR", "Object creation error")
)

const MaxEmailLength = 255

var ErrMaxEmailLengthExceeded = fmt.Errorf("email cannot contain more than %d characters", MaxEmailLength)
