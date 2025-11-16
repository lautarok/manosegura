package docs

import (
	"encoding/json"
	"go/ast"
	"strings"
)

func nameToPath(name string) string {
	return strings.Join(strings.Split(strings.TrimSuffix(name, "_controller.go"), "_"), "-")
}

func FindControllerDeclarations(fileDeclarations *FileDeclarations) {
	for controllerName, controller := range fileDeclarations.fileSet.controllers {
		for _, decl := range controller.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if fn.Doc != nil {
					controller := Controller{
						BasePath: nameToPath(controllerName),
						Returns:  map[string]string{},
						Examples: map[string]any{},
					}
					for _, comment := range fn.Doc.List {
						commentText := comment.Text
						if trimmed, found := strings.CutPrefix(commentText, "// @Router "); found {
							splitted := strings.Split(trimmed, " ")
							controller.Tags = []string{
								func() string {
									controllerName := strings.Join(strings.Split(controllerName, "_"), " ")
									controllerName = strings.ToUpper(string(controllerName[0])) + controllerName[1:]
									controllerName = strings.TrimSuffix(controllerName, "controller.go")
									return controllerName
								}(),
							}
							controller.Method = splitted[0]
							if len(splitted) >= 2 {
								controller.Path = splitted[1]
							} else {
								controller.Path = ""
							}
						} else if trimmed, found := strings.CutPrefix(commentText, "// @Description "); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							controller.Description = trimmed
						} else if trimmed, found := strings.CutPrefix(commentText, "// @Summary "); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							controller.Summary = trimmed
						} else if trimmed, found := strings.CutPrefix(commentText, "// @Accept "); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							controller.Accept = trimmed
						} else if trimmed, found := strings.CutPrefix(commentText, "// @Returns "); found {
							splitted := strings.Split(trimmed, " ")

							var example map[string]any
							if len(splitted) >= 3 {
								raw := strings.ReplaceAll(strings.Trim(strings.TrimPrefix(strings.Join(splitted[2:], " "), "example:"), "\""), "\\", "")
								json.Unmarshal([]byte(raw), &example)
							}

							if example != nil {
								controller.Examples[splitted[0]] = example
							}

							controller.Tags = []string{
								tagFromString(controllerName),
							}
							if len(splitted) >= 2 {
								controller.Returns[splitted[0]] = splitted[1]
							}
						} else if trimmed, found := strings.CutPrefix(commentText, "// @BodyRequest "); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							splitted := strings.Split(trimmed, " ")
							controller.Body = splitted[0]
						} else if trimmed, found := strings.CutPrefix(commentText, "// @QueryParams "); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							splitted := strings.Split(trimmed, " ")
							controller.QueryParams = splitted[0]
						} else if trimmed, found := strings.CutPrefix(commentText, "// @RouteParams "); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							splitted := strings.Split(trimmed, " ")
							controller.RouteParams = splitted[0]
						} else if _, found := strings.CutPrefix(commentText, "// @BearerAuth"); found {
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							controller.BearerAuth = true
						}
					}
					fileDeclarations.Controllers = append(fileDeclarations.Controllers, &controller)
				}
			}
		}
	}
}
