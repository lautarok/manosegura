package docs

import (
	"log"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/lautarok/manosegura/src/internal/exceptions"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	"github.com/lautarok/manosegura/src/internal/modules/users/domain"
)

func GetOpenApiSpec() *openapi3.T {

	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   "Mano Segura api documentation",
			Version: "1.0.0",
		},
		Paths: &openapi3.Paths{},
		Components: &openapi3.Components{
			Schemas: make(map[string]*openapi3.SchemaRef),
		},
	}

	userSchema, err := openapi3gen.NewSchemaRefForValue(&domain.User{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	paginationDtoSchema, err := openapi3gen.NewSchemaRefForValue(&dto.PaginationDto{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	appErrorSchema, err := openapi3gen.NewSchemaRefForValue(&exceptions.AppError{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	doc.Components.Schemas["User"] = userSchema
	doc.Components.Schemas["PaginationDto"] = paginationDtoSchema
	doc.Components.Schemas["AppError"] = appErrorSchema

	doc.Components.SecuritySchemes = openapi3.SecuritySchemes{
		"bearerAuth": &openapi3.SecuritySchemeRef{
			Value: &openapi3.SecurityScheme{
				Type:         "http",
				Scheme:       "Bearer",
				BearerFormat: "JWT",
			},
		},
	}

	return doc
}
