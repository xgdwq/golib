package main

import (
	"fmt"
	"format/flib"
)

func main() {

	var testMap = make(map[string]interface{})

	testMap["k1"] = 1
	testMap["k2"] = []string{"yuzhi", "ivy"}
	testMap["k3"] = map[string]string{
		"yuzhi": "wang",
		"ivy":   "liu",
		"eric":  "wang",
	}
	fmt.Println(flib.FormatJson(testMap))
}
