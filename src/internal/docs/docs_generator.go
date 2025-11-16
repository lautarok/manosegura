package docs

import (
	"strings"

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
	decls := NewFileDeclarations(fileSet)

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

	for _, domain := range decls.Domains {
		newSchema := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Title:      domain.Name,
				Properties: openapi3.Schemas{},
			},
		}

		for _, field := range domain.Fields {
			if field.Validation["required"] == "true" {
				newSchema.Value.Required = append(newSchema.Value.Required, field.Name)
			}
			newField := &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: &openapi3.Types{
						field.Type,
					},
				},
			}

			splittedFType := strings.Split(field.Type, ".")
			if pop := splittedFType; docsGenerator.oas.Components.Schemas[pop[len(splittedFType)-1]] != nil {
				newField.Ref = "#/components/schemas/" + pop[len(pop)-1]
			}

			if field.Example != "" {
				newField.Value.Example = field.Example
			}

			newSchema.Value.Properties[field.Name] = newField
		}

		docsGenerator.oas.Components.Schemas[domain.Name] = newSchema
	}

	for _, dto := range decls.Dtos {
		newSchema := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Title:      dto.Name,
				Properties: openapi3.Schemas{},
			},
		}

		for _, field := range dto.Fields {
			if field.Validation["required"] == "true" {
				newSchema.Value.Required = append(newSchema.Value.Required, field.Name)
			}
			newField := &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: &openapi3.Types{
						field.Type,
					},
				},
			}

			splittedFType := strings.Split(field.Type, ".")
			if pop := splittedFType; docsGenerator.oas.Components.Schemas[pop[len(splittedFType)-1]] != nil {
				newField.Ref = "#/components/schemas/" + pop[len(pop)-1]
			}

			if field.Example != "" {
				newField.Value.Example = field.Example
			}

			newSchema.Value.Properties[field.Name] = newField
		}

		docsGenerator.oas.Components.Schemas[dto.Name] = newSchema
	}

	for _, controller := range decls.Controllers {
		operation := openapi3.NewOperation()
		operation.Tags = controller.Tags
		operation.Description = controller.Description
		operation.Summary = controller.Summary

		for returnsKey, returnsValue := range controller.Returns {
			if operation.Responses == nil {
				operation.Responses = openapi3.NewResponses()
			}

			newResponse := openapi3.NewResponse()
			newResponse.Content = openapi3.NewContent()

			newResponse.Content[returnsKey] = &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{
					Ref: "#/components/schemas/" + returnsValue,
				},
			}

			operation.Responses.Set(returnsKey, &openapi3.ResponseRef{
				Value: newResponse,
			})
		}

		path := docsGenerator.oas.Paths.Find(controller.Path)
		if path == nil {
			path = &openapi3.PathItem{}
			docsGenerator.oas.Paths.Set(controller.Path, path)
		}

		switch controller.Method {
		case "get":
			path.Get = operation
		case "post":
			path.Post = operation
		case "put":
			path.Put = operation
		case "patch":
			path.Patch = operation
		case "delete":
			path.Delete = operation
		}
	}
	return nil
}

func (docsGenerator *DocsGenerator) Struct() *openapi3.T {
	return docsGenerator.oas
}
