package datautil

import (
	"encoding/base64"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func RandStringBytes(n int) string {
	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[seedRand.Intn(len(letterBytes)-1)]
	}
	return string(b)
}

func RandBytes(n int) []byte {
	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[seedRand.Intn(len(letterBytes)-1)]
	}
	return b
}

func EncodeByteToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodeStringToByte(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func XORSByte(a []byte, b []byte) []byte {
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

func XORString(a string, b string) string {
	return base64.StdEncoding.EncodeToString(XORSByte([]byte(a), []byte(b)))
}
