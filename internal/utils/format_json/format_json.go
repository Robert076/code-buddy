package format_json

import (
	"encoding/json"
)

func FormatNiceStringFromUglyJSON(body []byte) string {
	var parsed struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(body, &parsed); err == nil && parsed.Source != "" {
		return parsed.Source
	} else {
		return string(body)
	}
}
