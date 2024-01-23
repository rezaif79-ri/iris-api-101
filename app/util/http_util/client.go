package httputil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// HTTPCLient is a wrapper of default http client package
// To make ease of use with general function purpose
type HTTPClient struct {
	client *http.Client
}

type HTTPResponse struct {
	Status  int
	Message string
	Error   error
	Data    interface{}
}

// NewHTTPClient to inititate new net http client
// HTTPClient can be reuseable for different request
func NewHTTPClient(timeout time.Duration) HTTPClient {
	return HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetAPI require target uri and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) GetAPI(uri string, callback func(*http.Response, error)) {
	res, err := h.client.Get(uri)
	if err != nil {
		callback(nil, err)
		return
	}
	defer res.Body.Close()
	callback(res, err)
}

// PostAPI require target uri, body and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) PostAPI(uri string, body interface{}, callback func(*http.Response, error)) {
	bodyLoad, err := json.Marshal(body)
	res, err := h.client.Post(uri, "application/json", bytes.NewReader(bodyLoad))
	if err != nil {
		callback(nil, err)
		return
	}
	defer res.Body.Close()
	callback(res, err)
}

// PutAPI require target uri, body and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) PutAPI(uri string, body interface{}, callback func(*http.Response, error)) {
	bodyLoad, err := json.Marshal(body)
	res, err := h.client.Post(uri, "application/json", bytes.NewReader(bodyLoad))
	if err != nil {
		callback(nil, err)
		return
	}
	defer res.Body.Close()
	callback(res, err)
}

// DeleteAPI require target uri, body and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) DeleteAPI(uri string, body interface{}, callback func(*http.Response, error)) {
	bodyLoad, err := json.Marshal(body)
	res, err := h.client.Post(uri, "application/json", bytes.NewReader(bodyLoad))
	if err != nil {
		callback(nil, err)
		return
	}
	defer res.Body.Close()
	callback(res, err)
}
