package docs

import (
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

func FindDomainDeclarations(fileDeclarations *FileDeclarations) {
	for _, domain := range fileDeclarations.fileSet.domains {
		for _, decl := range domain.Decls {
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

				newDomain := Domain{}
				for _, spec := range decl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						newDomain.Name = typeSpec.Name.Name

						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							if structType.Fields == nil {
								continue
							}

							for _, typeData := range structType.Fields.List {
								if typeData.Tag == nil {
									continue
								}

								fType := "string"
								if fieldType, ok := typeData.Type.(*ast.Ident); ok {
									fType = fieldType.String()
								} else if fieldType, ok := typeData.Type.(*ast.StarExpr); ok {
									if fieldType2, ok := fieldType.X.(*ast.SelectorExpr); ok {
										if _, ok := fieldType2.X.(*ast.Ident); ok {
											fType = fieldType2.Sel.Name
										}
									} else if ident, ok := fieldType.X.(*ast.Ident); ok {
										fType = ident.Name
									}
								} else if fieldType, ok := typeData.Type.(*ast.SelectorExpr); ok {
									if ident, ok := fieldType.X.(*ast.Ident); ok {
										fType = ident.Name
									}
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

									newDomain.Fields = append(newDomain.Fields, &DomainField{
										Name:    tagJsonName,
										Type:    fType,
										Example: example,
									})
								}
							}
						}
					}

					fileDeclarations.Domains = append(fileDeclarations.Domains, &newDomain)

					break
				}
			}
		}
	}
}
