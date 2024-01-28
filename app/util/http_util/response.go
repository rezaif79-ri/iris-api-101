package httputil

import (
	"io"
)

func ConvertBodyToBytes(body io.ReadCloser) ([]byte, error) {
	return io.ReadAll(body)
}
