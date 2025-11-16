package docs

import (
	"log"
	"strconv"
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

func (docsGenerator *DocsGenerator) createDomains(decls []*Domain) {
	for _, domain := range decls {
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
}

func (docsGenerator *DocsGenerator) createDtos(decls []*Dto) {
	for _, dto := range decls {
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

			if field.Validation["min"] != "" {
				min, err := strconv.Atoi(field.Validation["min"])
				if err != nil {
					log.Fatal(err)
				}
				min2 := float64(min)
				newField.Value.Min = (&min2)
			}

			if field.Validation["max"] != "" {
				max, err := strconv.Atoi(field.Validation["max"])
				if err != nil {
					log.Fatal(err)
				}
				max2 := float64(max)
				newField.Value.Max = (&max2)
			}

			if field.Validation["gte"] != "" {
				gte, err := strconv.Atoi(field.Validation["gte"])
				if err != nil {
					log.Fatal(err)
				}
				gte2 := float64(gte)
				newField.Value.Min = (&gte2)
			}

			if field.Validation["lte"] != "" {
				lte, err := strconv.Atoi(field.Validation["lte"])
				if err != nil {
					log.Fatal(err)
				}
				lte2 := float64(lte)
				newField.Value.Max = (&lte2)
			}

			newSchema.Value.Properties[field.Name] = newField
		}

		docsGenerator.oas.Components.Schemas[dto.Name] = newSchema
	}

	docsGenerator.oas.Components.Schemas["Empty"] = &openapi3.SchemaRef{
		Value: &openapi3.Schema{},
	}
}

func (docsGenerator *DocsGenerator) createControllers(decls []*Controller) {
	for _, controller := range decls {
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

			var examplesMap openapi3.Examples

			for exampleKey, example := range controller.Examples {
				if exampleKey == returnsKey {
					if exampleContent, ok := example.(map[string]any); ok {
						if examplesMap == nil {
							examplesMap = openapi3.Examples{}
						}

						examplesMap[exampleKey] = &openapi3.ExampleRef{
							Value: &openapi3.Example{
								Value: exampleContent,
							},
						}
					}
				}
			}

			if value, found := strings.CutPrefix(returnsValue, "[]"); found {
				newResponse.Content["application/json"] = &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Items: &openapi3.SchemaRef{
								Ref: "#/components/schemas/" + value,
							},
						},
					},
					Examples: examplesMap,
				}
			} else {
				newResponse.Content["application/json"] = &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Ref: "#/components/schemas/" + returnsValue,
					},
					Examples: examplesMap,
				}
			}

			operation.Responses.Set(returnsKey, &openapi3.ResponseRef{
				Value: newResponse,
			})
		}

		if controller.Body != "" {
			operation.RequestBody = &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: "#/components/schemas/" + controller.Body,
							},
						},
					},
				},
			}
		}

		parameters := []*openapi3.ParameterRef{}

		if controller.QueryParams != "" {
			schema := docsGenerator.oas.Components.Schemas[controller.QueryParams].Value
			parameters = append(parameters, SchemaAsQueryParams(schema)...)
		}

		if controller.RouteParams != "" {
			schema := docsGenerator.oas.Components.Schemas[controller.RouteParams].Value
			parameters = append(parameters, SchemaAsRouteParams(schema)...)
		}

		if controller.BearerAuth {
			operation.Security = &openapi3.SecurityRequirements{
				openapi3.SecurityRequirement{
					"bearerAuth": []string{},
				},
			}
		}

		schemaParameters := openapi3.NewParameters()
		schemaParameters = append(schemaParameters, parameters...)

		operation.Parameters = schemaParameters

		path := docsGenerator.oas.Paths.Find("/" + controller.BasePath + controller.Path)
		if path == nil {
			path = &openapi3.PathItem{}
			docsGenerator.oas.Paths.Set("/"+controller.BasePath+controller.Path, path)
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
			SecuritySchemes: openapi3.SecuritySchemes{
				"bearerAuth": &openapi3.SecuritySchemeRef{
					Value: &openapi3.SecurityScheme{
						Type:         "http",
						Scheme:       "bearer",
						BearerFormat: "JWT",
					},
				},
			},
		},
	}

	docsGenerator.createDomains(decls.Domains)
	docsGenerator.createDomains(decls.Domains)

	docsGenerator.createDtos(decls.Dtos)
	docsGenerator.createDtos(decls.Dtos)

	docsGenerator.createControllers(decls.Controllers)

	return nil
}

func (docsGenerator *DocsGenerator) Struct() *openapi3.T {
	return docsGenerator.oas
}
