package docs

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

type FileSet struct {
	controllers map[string]*ast.File
}

func NewFileSet() *FileSet {
	newFileSet := FileSet{
		make(map[string]*ast.File),
	}
	err := newFileSet.FetchFiles()
	if err != nil {
		log.Fatal(err)
	}
	return &newFileSet
}

func (fileSet *FileSet) FetchFiles() error {
	baseDir := "./src/internal/modules"
	modules, err := os.ReadDir(baseDir)
	if err != nil {
		return err
	}

	tokenFileSet := token.NewFileSet()

	for _, module := range modules {
		if !module.IsDir() {
			continue
		}

		subfolders, err := os.ReadDir(baseDir + "/" + module.Name())
		if err != nil {
			return err
		}

		for _, subfolder := range subfolders {
			if !module.IsDir() {
				continue
			}

			if subfolder.Name() == "controllers" {
				controllers, err := os.ReadDir(baseDir + "/" + module.Name() + "/" + subfolder.Name())
				if err != nil {
					return err
				}

				for _, controller := range controllers {
					if controller.IsDir() || filepath.Ext(controller.Name()) != ".go" {
						continue
					}

					file, err := parser.ParseFile(tokenFileSet, baseDir+"/"+module.Name()+"/"+subfolder.Name()+"/"+controller.Name(), nil, parser.ParseComments)
					if err != nil {
						return err
					}

					fileSet.controllers[controller.Name()] = file
				}
			}
		}
	}

	return nil
}
