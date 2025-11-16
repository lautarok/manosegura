package docs

import "github.com/getkin/kin-openapi/openapi3"

func SchemaAsQueryParams(schema *openapi3.Schema) []*openapi3.ParameterRef {
	params := []*openapi3.ParameterRef{}

	for propName, propSchema := range schema.Properties {
		isRequired := false
		for _, required := range schema.Required {
			if required == propName {
				isRequired = true
			}
		}

		param := &openapi3.Parameter{
			In:       "query",
			Name:     propName,
			Required: isRequired,
			Schema:   propSchema,
		}

		params = append(params, &openapi3.ParameterRef{
			Value: param,
		})
	}

	return params
}

func SchemaAsRouteParams(schema *openapi3.Schema) []*openapi3.ParameterRef {
	params := []*openapi3.ParameterRef{}

	for propName, propSchema := range schema.Properties {
		isRequired := false
		for _, required := range schema.Required {
			if required == propName {
				isRequired = true
			}
		}

		param := &openapi3.Parameter{
			In:       "path",
			Name:     propName,
			Required: isRequired,
			Schema:   propSchema,
		}

		params = append(params, &openapi3.ParameterRef{
			Value: param,
		})
	}

	return params
}
