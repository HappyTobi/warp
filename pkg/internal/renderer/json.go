package renderer

import (
	"bytes"
	"encoding/json"
)

func JsonInterface(data interface{}) (string, error) {
	dat, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return JsonByte(dat)
}

/*
Print out bytes array as prettyJson
*/
func JsonByte(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "\t"); err != nil {
		return "", err
	}

	return string(prettyJSON.Bytes()), nil
}
