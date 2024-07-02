package URLEncoder

import (
	"encoding/json"
	"fmt"
	"strings"
)

func EncodeBody(fields string) (string, error) {

	var postData map[string]interface{}

	err := json.Unmarshal([]byte(fields), &postData)
	if err != nil {
		return "", err
	}

	var keyPairs []string

	for k, v := range postData {
		keyPair := fmt.Sprintf("%s=%v", k, v)
		keyPairs = append(keyPairs, keyPair)
	}

	postBody := strings.Join(keyPairs, "&")
	return postBody, nil
}
