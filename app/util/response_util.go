package util

import "github.com/kataras/iris/v12/context"

type MapString map[string]interface{}

func RestWrapperObject(status int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

// IrisJSONResponse will write response to client by iris context
func IrisJSONResponse(ctx *context.Context, statusCode int, message string, data interface{}) {
	ctx.StatusCode(statusCode)
	ctx.JSON(RestWrapperObject(statusCode, message, data))
	return
}
