package utils

import (
	"bytes"
	"net/url"
)

func EncodeKey(prefix, suffix string, keys ...string) string {
	var buf bytes.Buffer
	buf.WriteString(prefix)
	buf.WriteString(suffix)
	for _, v := range keys {
		if v != "" {
			buf.WriteString(":")
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
