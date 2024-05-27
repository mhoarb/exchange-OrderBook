package app

import (
	"encoding/json"
	"fmt"
)

func JsonPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	return string(s)
}
