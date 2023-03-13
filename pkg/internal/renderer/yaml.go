package renderer

import (
	"gopkg.in/yaml.v3"
)

func (yr *yamlRenderer) Render(data interface{}) (string, error) {
	dat, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(dat), err
}

/*
Print out bytes array as Yaml
*/
func (yr *yamlRenderer) RenderBytes(data []byte) (string, error) {
	var generic map[string]interface{}
	if err := yaml.Unmarshal(data, &generic); err != nil {
		return "", err
	}

	return yr.Render(generic)
}
