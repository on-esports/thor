package thor

import (
	"context"
	"net/http"
)

// Requester is a http requester logic handler. It is responsible for handling all the http requests
type Requester interface {
	GET(ctx context.Context, response any, opt HttpOptions) (statusCode int, err error)
	POST(ctx context.Context, opt HttpOptions) (statusCode int, resp []byte, err error)
	PUT(ctx context.Context, opt HttpOptions) (statusCode int, resp []byte, err error)
	DELETE(ctx context.Context, opt HttpOptions) (statusCode int, resp []byte, err error)
}

// api is the struct that handles making HTTP requests.
type api struct {
	request *http.Client
}

// New creates a new api instance with a default HTTP client.
func New() Requester {
	return &api{
		request: &http.Client{},
	}
}
