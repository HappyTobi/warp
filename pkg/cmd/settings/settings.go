package settings

import (
	_ "embed"
	"os"
)

//go:embed logo.png
var logo []byte

func StoreImage(path string) error {
	return os.WriteFile(path, logo, 0644)
}
