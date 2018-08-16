package iutil

import (
	"encoding/json"
)

// ToJSONString 将任意类型的数据，转成json格式的字符串
func ToJSONString(s interface{}) (string, error) {
	byt, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
