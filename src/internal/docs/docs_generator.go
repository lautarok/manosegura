package docs

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type DocsGenerator struct {
	oas         *openapi3.T
	title       string
	description string
	version     string
}

type DocsGeneratorConfig struct {
	Title       string
	Description string
	Version     string
}

func NewDocsGenerator(config *DocsGeneratorConfig) *DocsGenerator {
	return &DocsGenerator{
		oas:         nil,
		title:       config.Title,
		description: config.Description,
		version:     config.Version,
	}
}

func (docsGenerator *DocsGenerator) GenerateOas() error {
	fileSet := NewFileSet()
	declarations := NewFileDeclarations(fileSet)

	docsGenerator.oas = &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       docsGenerator.title,
			Description: docsGenerator.description,
			Version:     docsGenerator.version,
		},
		Paths: &openapi3.Paths{},
		Components: &openapi3.Components{
			Schemas: make(map[string]*openapi3.SchemaRef),
		},
	}

	for _, declaration := range declarations.FileDeclarations {
		operation := openapi3.NewOperation()

		if declaration.Value["method"] == "get" {
			docsGenerator.oas.Paths.Set(declaration.Value["path"], &openapi3.PathItem{
				Get: operation,
			})
		}
	}
	return nil
}

func (docsGenerator *DocsGenerator) Struct() *openapi3.T {
	return docsGenerator.oas
}
