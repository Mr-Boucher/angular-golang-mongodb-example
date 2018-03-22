package utils

import (
	"bytes"
	"fmt"
)

//
func CreateKeyValuePairs(m map[string]string, defaultValue string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}

	result := defaultValue
	if b.Len() > 0 {
		result = b.String()
	}
	return result
}
