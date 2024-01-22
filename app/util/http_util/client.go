package httputil

import (
	"net/http"
	"time"
)

type HTTPClient http.Client

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		Timeout: timeout,
	}
}
