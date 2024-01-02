package util

import "os"

func GetEnv(key string, defValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defValue
}
