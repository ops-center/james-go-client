package james

import (
	"encoding/json"
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

func unmarshalToJamesServerError(r *http.Response) *JamesServerError {
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return &JamesServerError{
			StatusCode: 500,
			Type:       "InternalServerError",
		}
	}
	je := JamesServerError{}
	if err := json.Unmarshal(bodyData, &je); err != nil {
		return &JamesServerError{
			StatusCode: 500,
			Type:       "InternalServerError",
		}
	}

	if je.StatusCode == 0 {
		je.StatusCode = r.StatusCode
	}

	return &je
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
