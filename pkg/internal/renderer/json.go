package renderer

import (
	"bytes"
	"encoding/json"
)

/*
Print out bytes array as prettyJson
*/
func PrettyJson(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "\t"); err != nil {
		return "", err
	}

	return string(prettyJSON.Bytes()), nil
}
