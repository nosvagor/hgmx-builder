package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Ref is a helper function to return a pointer to a value.
//
// Example:
//
//	needsRef(utils.Ref(&v))
//
// is the same as:
//
//	v := "value"
//	needsRef(&v)
func Ref[T any](v T) *T {
	return &v
}

// If is a ternary operator for Go.
//
// Example:
//
//	utils.If(condition, trueVal, falseVal)
//
// is the same as:
//
//	if condition {
//		return trueVal
//	}
//	return falseVal
func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// JSONPretty returns a indented JSON encoding of v.
//
// Example:
//
//	utils.JSONPretty(map[string]string{"key": "value"})
//
// is the same as:
//
//	json.MarshalIndent(map[string]string{"key": "value"}, "", "  ")
//
// Returns formatted JSON string
func JSONPretty(v any) string {
	switch data := v.(type) {
	case []byte:
		var buf bytes.Buffer
		err := json.Indent(&buf, data, "", "  ")
		if err != nil {
			return ""
		}
		return buf.String()
	case string:
		var buf bytes.Buffer
		err := json.Indent(&buf, []byte(data), "", "  ")
		if err != nil {
			return ""
		}
		return buf.String()
	default:
		jsonBytes, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return ""
		}
		return string(jsonBytes)
	}
}

func Debug(v any, prefix ...string) {
	prefixStr := ""
	if len(prefix) > 0 {
		prefixStr = prefix[0] + " "
	}
	fmt.Println(prefixStr, JSONPretty(v))
}
