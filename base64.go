package iutil

import (
	"encoding/base64"
	"strings"
)

// Base64URLDecode base64 decode 安全版
func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)

	return base64.URLEncoding.DecodeString(data)
}

// Base64UrlSafeEncode base64 encode 安全版
func Base64UrlSafeEncode(source []byte) string {
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}
