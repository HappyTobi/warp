package renderer

import (
	"fmt"
	"os"
)

type Renderer interface {
	Render(data interface{}) (string, error)
	RenderBytes(data []byte) (string, error)
}

func NewRenderer(output string) Renderer {
	switch output {
	case "json":
		return &jsonRenderer{}
	case "yaml":
		return &yamlRenderer{}
		// case "csv":
		// 	return &csvRenderer{}
	}

	fmt.Printf("unknown output format: %s", output)
	os.Exit(1)
	return nil
}
