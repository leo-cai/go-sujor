package utils

import (
	"encoding/base64"
)

//const KEY_TEXT  = "sujor_api20171116?leo"

// 加密
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// 解密
func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
