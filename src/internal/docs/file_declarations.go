package docs

import (
	"go/ast"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FileDeclarations struct {
	FileDeclarations []*Declaration
	fileSet          *FileSet
}

const (
	PathDeclaration = iota
)

type Declaration struct {
	Name  string
	Type  int
	Value map[string]string
}

func NewFileDeclarations(fileSet *FileSet) *FileDeclarations {
	fd := &FileDeclarations{
		fileSet:          fileSet,
		FileDeclarations: []*Declaration{},
	}
	fd.FetchDeclarations()
	return fd
}

func (fileDeclarations *FileDeclarations) FetchDeclarations() {
	for controllerName, controller := range fileDeclarations.fileSet.controllers {
		for _, decl := range controller.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if fn.Doc != nil {
					for _, comment := range fn.Doc.List {
						commentText := comment.Text
						if trimmed, found := strings.CutPrefix(commentText, "// @Router "); found {
							splitted := strings.Split(trimmed, " ")
							fileDeclarations.FileDeclarations = append(fileDeclarations.FileDeclarations, &Declaration{
								Name: func() string {
									trimmed := strings.TrimSuffix(controllerName, "_controller.go")
									splitted := strings.Split(trimmed, "_")
									name := strings.Join(splitted, " ")
									c := cases.Title(language.English)
									return c.String(name)
								}(),
								Type: PathDeclaration,
								Value: map[string]string{
									"method": splitted[0],
									"path":   splitted[1],
								},
							})
						}
					}
				}
			}
		}
	}
}
