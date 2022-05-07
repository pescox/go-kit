package exp

import (
	"encoding/json"
	"os"
	"unicode/utf8"

	"github.com/pescox/go-kit/log"
)

func BytesCount(s string) int {
	return len(s)
}

func CharsCount(s string) int {
	return utf8.RuneCountInString(s)
}

func GetEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

// ShortenString shorten string to length.
func ShortenString(v string, n int) string {
	if len(v) <= n {
		return v
	}
	return v[:n]
}

// JsonMarshal marshal object to json string.
func JsonMarshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.ErrorF("json marshal err: %s", err.Error())
		return ""
	}
	return string(bytes)
}

// JsonUnmarshal unmarshal json string to object.
func JsonUnmarshal(str string, v interface{}) {
	if err := json.Unmarshal([]byte(str), &v); err != nil {
		log.ErrorF("json unmarshal err: %s", err.Error())
	}
}
