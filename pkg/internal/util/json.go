package util

import "encoding/json"

func DeserializeJsonData(data []byte, dest interface{}) error {
	return json.Unmarshal(data, dest)
}
