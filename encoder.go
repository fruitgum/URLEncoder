package urlencoder

import (
	"fmt"
	"strings"
)

func EncodeBody(body map[string]interface{}) (string, error) {

	var keyPairs []string

	for k, v := range body {
		keyPair := fmt.Sprintf("%s=%v", k, v)
		keyPairs = append(keyPairs, keyPair)
	}

	postBody := strings.Join(keyPairs, "&")
	return postBody, nil
}
