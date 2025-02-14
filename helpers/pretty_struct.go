package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PrintData(prefix string, data any) {
	val, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println(strings.TrimSpace(prefix) + " : " + string(val))
}

func PrintDataIndent(prefix string, data any) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(strings.TrimSpace(prefix) + " : " + string(val))
}

func ToString(j any) string {
	indent, _ := json.MarshalIndent(j, "", "  ")
	return string(indent)
}
