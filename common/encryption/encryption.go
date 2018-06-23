package encryption

import (
	"crypto/sha256"
	"fmt"
	"encoding/base64"
)

func Gensha256(phone string, nowtime int64, salt string) []byte {
	s := fmt.Sprintf("%s%d%s", phone, nowtime, salt)

	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}

func GenBase64(input string)  []byte {
	s:=base64.StdEncoding.EncodeToString([]byte(input))
	return []byte(s)
}

