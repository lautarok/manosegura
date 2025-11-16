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
	dtos        map[string]*ast.File
	domains     map[string]*ast.File
}

func NewFileSet() *FileSet {
	newFileSet := FileSet{
		controllers: make(map[string]*ast.File),
		dtos:        make(map[string]*ast.File),
		domains:     make(map[string]*ast.File),
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

			if subfolder.Name() == "controllers" || subfolder.Name() == "domain" || subfolder.Name() == "dto" {
				sets, err := os.ReadDir(baseDir + "/" + module.Name() + "/" + subfolder.Name())
				if err != nil {
					return err
				}

				for _, set := range sets {
					if set.IsDir() || filepath.Ext(set.Name()) != ".go" {
						continue
					}

					file, err := parser.ParseFile(tokenFileSet, baseDir+"/"+module.Name()+"/"+subfolder.Name()+"/"+set.Name(), nil, parser.ParseComments)
					if err != nil {
						return err
					}

					if subfolder.Name() == "controllers" {
						fileSet.controllers[set.Name()] = file
					} else if subfolder.Name() == "dto" {
						fileSet.dtos[set.Name()] = file
					} else if subfolder.Name() == "domain" {
						fileSet.domains[set.Name()] = file
					}
				}
			}
		}
	}

	return nil
}
