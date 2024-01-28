package httputil

import (
	"io"
)

type HTTPResponse struct {
	Status  int
	Message string
	Error   error
	Data    []byte
}

func convertBodyToBytes(body io.ReadCloser) ([]byte, error) {
	return io.ReadAll(body)
}
