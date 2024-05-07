package dhjson

import (
	"encoding/json"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	// "fmt"
	// "strconv"
)

// JsonEncode
// @Description: 结构体转json
// @param v
// @return string
func JsonEncode(v interface{}) string {
	encodeData, err := json.Marshal(v)
	if nil != err {
		dhlog.Info("JsonEncode Error: ", err)
		return ""
	}
	return string(encodeData)
}

// JsonEncodeIndent
// @Description: 结构体转json，待pretty效果
// @param v
// @return string
func JsonEncodeIndent(v interface{}) string {
	encodeData, err := json.MarshalIndent(v, "", "    ")
	if nil != err {
		dhlog.Info("JsonEncode Error: ", err)
		return ""
	}
	return string(encodeData)
}

func PrintJsonEncodeIndent(v interface{}) {
	encodeData, err := json.MarshalIndent(v, "", "    ")
	if nil != err {
		dhlog.Warn("JsonEncode Error: ", err)
	}
	dhlog.Info(string(encodeData))
}

func Json2Struct(jsonStr string, v any) error {
	return json.Unmarshal([]byte(jsonStr), v)
}
