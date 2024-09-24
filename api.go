package thor

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HttpOptions needs to make generic http calls to the third api with options
type HttpOptions struct {
	Url     string            // Provide full url here
	Headers map[string]string // Request headers
	Body    any               // Request body, for POST, PUT, DELETE methods
	Timeout time.Duration     // Optional timeout for the request
}

// GET sends a GET request to the specified Endpoint with options.
// for response parameter provide anything that can be marshalled with json
func (a *api) GET(ctx context.Context, response any, opt HttpOptions) (int, error) {
	const action = "thor.GET method: %w"
	// Create the request
	var err error
	opt.Url, err = a.validateUrl(opt.Url)
	if err != nil {
		return 0, fmt.Errorf(action, err)
	}

	req, err := a.createRequest(ctx, http.MethodGet, opt)
	if err != nil {
		return 0, fmt.Errorf(action, err)
	}
	// Execute the request
	statusCode, respBody, err := a.executeRequest(req)
	if err != nil {
		return statusCode, fmt.Errorf(action, err)
	}

	// Unmarshal the response if needed
	if response != nil {
		if err = a.unmarshalResponse(respBody, response); err != nil {
			return statusCode, fmt.Errorf(action, err)
		}
	}
	return statusCode, nil
}

// POST sends a POST request to the specified Endpoint with options.
func (a *api) POST(ctx context.Context, opt HttpOptions) (statusCode int, resp []byte, err error) {
	const action = "thor.POST method: %w"
	// Create the request
	opt.Url, err = a.validateUrl(opt.Url)
	if err != nil {
		err = fmt.Errorf(action, err)
		return
	}
	req, err := a.createRequest(ctx, http.MethodPost, opt)
	if err != nil {
		err = fmt.Errorf(action, err)
		return
	}
	// Execute the request
	statusCode, resp, err = a.executeRequest(req)
	if err != nil {
		err = fmt.Errorf(action, err)
	}
	return
}

// PUT sends a PUT request to the specified Endpoint with options.
func (a *api) PUT(ctx context.Context, opt HttpOptions) (statusCode int, resp []byte, err error) {
	const action = "thor.PUT method: %w"
	// Create the request
	opt.Url, err = a.validateUrl(opt.Url)
	if err != nil {
		err = fmt.Errorf(action, err)
		return
	}
	req, err := a.createRequest(ctx, http.MethodPut, opt)
	if err != nil {
		return 0, nil, fmt.Errorf(action, err)
	}
	// Execute the request
	statusCode, resp, err = a.executeRequest(req)
	if err != nil {
		err = fmt.Errorf(action, err)
	}
	return
}

// DELETE sends a DELETE request to the specified Endpoint with options.
func (a *api) DELETE(ctx context.Context, opt HttpOptions) (statusCode int, resp []byte, err error) {
	const action = "thor.DELETE method: %w"
	// Create the request
	opt.Url, err = a.validateUrl(opt.Url)
	if err != nil {
		err = fmt.Errorf(action, err)
		return
	}
	req, err := a.createRequest(ctx, http.MethodDelete, opt)
	if err != nil {
		return 0, nil, fmt.Errorf(action, err)
	}
	// Execute the request
	statusCode, resp, err = a.executeRequest(req)
	if err != nil {
		err = fmt.Errorf(action, err)
	}
	return
}
