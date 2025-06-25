package inbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"git.sr.ht/~rockorager/go-jmap"
)

type JamesServerError struct {
	StatusCode int    `json:"statusCode"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	Details    string `json:"details"`
}

func newServerError(resp *http.Response, respErr error) *JamesServerError {
	if resp == nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("response is nil. err: %v", respErr),
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("failed to read response body: %v. error: %v", err, respErr),
		}
	}

	if !json.Valid(body) {
		return &JamesServerError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("response: %v. err: %v", string(body), respErr),
		}
	}

	var jamesServerError JamesServerError
	if err = json.Unmarshal(body, &jamesServerError); err != nil {
		return &JamesServerError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("failed to unmarshal error response: %v. error: %v", err, respErr),
		}
	}

	return &jamesServerError
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
		err = fmt.Errorf("%w : %s", err, j.Message)
	}
	if !isEmptyString(j.Details) {
		err = fmt.Errorf("%w : %s", err, j.Details)
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

var ErrMaxEmailLengthExceeded = fmt.Errorf("email address cannot consist of more than %d characters", MaxEmailLength)
