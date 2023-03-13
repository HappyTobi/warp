package renderer

import (
	"bytes"
	"encoding/json"
)

func (jr *jsonRenderer) Render(data interface{}) (string, error) {
	dat, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return jr.RenderBytes(dat)
}

/*
Print out bytes array as prettyJson
*/
func (jr *jsonRenderer) RenderBytes(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "\t"); err != nil {
		return "", err
	}

	return string(prettyJSON.Bytes()), nil
}
