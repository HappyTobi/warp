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
	case "yaml", "yml":
		return &yamlRenderer{}
	}
	fmt.Printf("unknown output format: %s", output)
	os.Exit(1)
	return nil
}

func IsRenderer(output string) bool {
	return output == "json" || output == "yaml" || output == "yml"
}

func NewCsvRenderer(settings *CsvSettings) *csvRenderer {
	return &csvRenderer{settings: settings}
}

func NewPdfRenderer(settings *PdfSettings) *pdfRenderer {
	return &pdfRenderer{settings: settings}
}
