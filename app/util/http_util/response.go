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

var mapResponseStatusMessage map[int]string = map[int]string{
	200: "http util: success",
	201: "http util: success",
	202: "http util: success",

	400: "http util: bad request",
	401: "http util: unauthorized",
	402: "http util: payment required",
	403: "http util: forbidden",
	404: "http util: not found",

	500: "http util: internal server error",
	501: "http util: unimplemented",
	502: "http util: bad gateway",
}

func getDefaultStatusMessage(code int) string {
	if v, ok := mapResponseStatusMessage[code]; ok {
		return v
	}
	return "http util: success"
}
