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
func (h *HTTPClient) GetAPI(uri string, callback func(*HTTPResponse)) {
	res, err := h.client.Get(uri)
	if err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    nil,
		})
		return
	}
	defer res.Body.Close()
	if body, err := convertBodyToBytes(res.Body); err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: "http_util: body reader error",
			Error:   err,
			Data:    nil,
		})
	} else {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    body,
		})
	}
}

// PostAPI require target uri, body and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) PostAPI(uri string, body interface{}, callback func(*HTTPResponse)) {
	bodyLoad, err := json.Marshal(body)
	res, err := h.client.Post(uri, "application/json", bytes.NewReader(bodyLoad))
	if err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    nil,
		})
		return
	}
	defer res.Body.Close()
	if body, err := convertBodyToBytes(res.Body); err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: "http_util: body reader error",
			Error:   err,
			Data:    nil,
		})
	} else {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    body,
		})
	}
}

// PutAPI require target uri, body and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) PutAPI(uri string, body interface{}, callback func(*HTTPResponse)) {
	bodyLoad, err := json.Marshal(body)
	res, err := h.client.Post(uri, "application/json", bytes.NewReader(bodyLoad))
	if err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: "http_util: request failed",
			Error:   err,
			Data:    nil,
		})
		return
	}
	defer res.Body.Close()
	if body, err := convertBodyToBytes(res.Body); err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: "http_util: body reader error",
			Error:   err,
			Data:    nil,
		})
	} else {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    body,
		})
	}
}

// DeleteAPI require target uri, body and callback as a function to retrieve repsonse or error
// Run NewRequest without headers and GET method
func (h *HTTPClient) DeleteAPI(uri string, body interface{}, callback func(*HTTPResponse)) {
	bodyLoad, err := json.Marshal(body)
	res, err := h.client.Post(uri, "application/json", bytes.NewReader(bodyLoad))
	if err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    nil,
		})
		return
	}
	defer res.Body.Close()
	if body, err := convertBodyToBytes(res.Body); err != nil {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: "http_util: body reader error",
			Error:   err,
			Data:    nil,
		})
	} else {
		callback(&HTTPResponse{
			Status:  res.StatusCode,
			Message: getDefaultStatusMessage(res.StatusCode),
			Error:   err,
			Data:    body,
		})
	}
}
