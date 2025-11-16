package docs

import (
	"go/ast"
	"strings"
)

func FindControllerDeclarations(fileDeclarations *FileDeclarations) {
	for controllerName, controller := range fileDeclarations.fileSet.controllers {
		for _, decl := range controller.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if fn.Doc != nil {
					controller := Controller{
						Returns: map[string]string{},
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
							controller.Path = splitted[1]
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
							controller.Tags = []string{
								tagFromString(controllerName),
							}
							splitted := strings.Split(trimmed, " ")
							if len(splitted) >= 2 {
								controller.Returns[splitted[0]] = splitted[1]
							}
						}
					}
					fileDeclarations.Controllers = append(fileDeclarations.Controllers, &controller)
				}
			}
		}
	}
}
