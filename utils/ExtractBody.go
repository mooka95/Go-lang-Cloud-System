package utils

import (
	"encoding/json"
	"io"
)

func ExtractBodyFromRequest(requestBody io.Reader) (map[string]interface{}, error) {
	var data map[string]interface{}
	body, err := io.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	//   var requestBody map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}
