package schemas

import (
	"embed"
)

func main() {
}

//go:embed schemas/*.json
var files embed.FS

func GetSchema(name string) ([]byte, error) {
	return files.ReadFile("schemas/" + name)
}

func GetMetadata(name string) ([]byte, error) {
	return files.ReadFile("metadata/" + name)
}
