package util

type MapString map[string]interface{}

func RestWrapperObject(status int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
