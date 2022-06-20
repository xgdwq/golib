package flib

import "encoding/json"

//json格式化
func FormatJson(i interface{}) string {
	bts, _ := json.MarshalIndent(i, "", "  ")
	return string(bts)
}
