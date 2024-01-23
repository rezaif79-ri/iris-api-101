package httputil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

type HTTPResponse struct {
}

func NewHTTPClient(timeout time.Duration) HTTPClient {
	return HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *HTTPClient) GetAPI(uri string, callback func(*http.Response, error)) {
	res, err := h.client.Get(uri)
	if err != nil {
		callback(nil, err)
		return
	}
	defer res.Body.Close()
	callback(res, err)
}

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
