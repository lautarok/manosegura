package docs

import "github.com/getkin/kin-openapi/openapi3"

func PaginationParamsDocs() openapi3.Parameters {
	return openapi3.Parameters{
		&openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name: "limit",
				In:   "query",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{
							openapi3.TypeInteger,
						},
						Format: "int32",
						Min: func() *float64 {
							n := float64(1)
							return &n
						}(),
						Max: func() *float64 {
							n := float64(25)
							return &n
						}(),
					},
				},
			},
		},
		&openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name: "page",
				In:   "query",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{
							openapi3.TypeInteger,
						},
						Format: "int32",
					},
				},
			},
		},
	}
}
