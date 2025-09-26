package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

func EmToHashSt(st string) string {
	data := []byte(st)
	hasher := sha1.New()
	hasher.Write(data)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
