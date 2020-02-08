package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

const apiEndpoint string = "https://pixe.la/v1"

func generateRequest(method string, path string, token *string, reqParams interface{}) (*http.Request, error) {
	var body io.Reader
	if reqParams != nil {
		params, err := json.Marshal(reqParams)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		body = bytes.NewBuffer(params)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", apiEndpoint, path), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != nil {
		req.Header.Set("X-USER-TOKEN", *token)
	}

	return req, nil
}