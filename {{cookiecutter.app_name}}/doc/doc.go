package doc

import (
	"github.com/jiaxinwang/common/fs"

	"github.com/swaggo/swag"
)

var doc = ""

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin API",
	Description: "An example of gin",
}

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}

func Init() {
	content, _ := fs.ReadBytes("__PATH_TO_SWAGGER_FILE__")
	doc = string(content)
	swag.Register(swag.Name, &s{})
}

func init() {
}
