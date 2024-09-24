package thor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func (a *api) validateUrl(baseUrl string) (string, error) {
	parse, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	if parse.Scheme == "" || parse.Host == "" {
		return "", errors.New("invalid url")
	}
	return parse.String(), nil
}

func (a *api) createRequest(ctx context.Context, method string, opt HttpOptions) (*http.Request, error) {
	// Prepare the request Endpoint with optional query parameters
	var bodyBytes []byte
	var err error
	if opt.Body != nil {
		bodyBytes, err = json.Marshal(opt.Body) // Marshal the body if it's not nil
		if err != nil {
			return nil, err
		}
	}
	// Create the request with the context, method, Endpoint, and body
	req, err := http.NewRequestWithContext(ctx, method, opt.Url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	// Set the authorization header if provided
	// Add any additional headers
	for key, value := range opt.Headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

// executeRequest performs the HTTP request and returns the status code and response body.
func (a *api) executeRequest(req *http.Request) (int, []byte, error) {
	// Perform the request
	resp, err := a.request.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, respBody, nil
}

// unmarshalResponse unmarshal the JSON response body into the provided response object.
func (a *api) unmarshalResponse(respBody []byte, response any) error {
	return json.Unmarshal(respBody, response)
}
