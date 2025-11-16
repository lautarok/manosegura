package docs

import (
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

func FindDtoDeclarations(fileDeclarations *FileDeclarations) {
	for _, dto := range fileDeclarations.fileSet.dtos {
		for _, decl := range dto.Decls {
			if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.TYPE {
				if decl.Doc != nil {
					hide := false
					for _, comment := range decl.Doc.List {
						if comment.Text == "// @Hide" {
							hide = true
							break
						}
					}
					if hide {
						continue
					}
				}

				newDto := Dto{}
				for _, spec := range decl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						newDto.Name = typeSpec.Name.Name

						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							if structType.Fields == nil {
								continue
							}

							for _, typeData := range structType.Fields.List {
								if typeData.Tag == nil {
									continue
								}

								fType := ResolveAstType(typeData.Type)

								switch fType {
								case "uuid.UUID":
									fType = "string"
								case "time.Time":
									fType = "date"
								}

								tagValue := typeData.Tag.Value
								tagStruct := reflect.StructTag(
									strings.Trim(tagValue, "`"),
								)

								tagJsonName := strings.Split(tagStruct.Get("json"), ",")[0]
								if tagJsonName != "-" {
									fValidation := tagStruct.Get("validate")
									validationMap := map[string]string{}
									for _, validation := range strings.Split(fValidation, ",") {
										splitted := strings.Split(validation, "=")
										validationKey := splitted[0]
										validationValue := "true"
										if len(splitted) == 2 {
											validationValue = splitted[1]
										}
										validationMap[validationKey] = validationValue
									}

									example := tagStruct.Get("example")

									if tagJsonName == "" {
										tagJsonName = typeData.Names[0].String()
									}

									newDto.Fields = append(newDto.Fields, &DtoField{
										Name:       tagJsonName,
										Type:       fType,
										Example:    example,
										Validation: validationMap,
									})
								}
							}
						}
					}

					fileDeclarations.Dtos = append(fileDeclarations.Dtos, &newDto)

					break
				}
			}
		}
	}
}
