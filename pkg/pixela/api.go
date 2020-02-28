package pixela

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

const apiEndpoint string = "https://pixe.la/v1"

func generateRequest(method string, url *url.URL, token *string, reqParams interface{}) (*http.Request, error) {
	var body io.Reader
	if reqParams != nil {
		params, err := json.Marshal(reqParams)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		body = bytes.NewBuffer(params)
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != nil {
		req.Header.Set("X-USER-TOKEN", *token)
	}

	return req, nil
}

func GenerateUrl(paths ...string) *url.URL {
	u, err := url.Parse(apiEndpoint)
	if err != nil {
		panic("the definition of `apiEndpoint` is wrong")
	}
	for _, p := range paths {
		u.Path = path.Join(u.Path, p)
	}
	return u
}
