package sha256util

import "crypto/sha256"

func SHA256GenSaltedHash(src []byte, salt []byte) []byte {
	saltSrc := xorByte(src, salt)
	hashByte := sha256.Sum256(saltSrc)
	return hashByte[:]
}

func xorByte(a []byte, b []byte) []byte {
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}
